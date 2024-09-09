package server

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
	"os/signal"
	"syscall"
	opapp "test-openapi/generated/openapi/app"
	"test-openapi/handlers"
	"test-openapi/middlewares"
	"time"
)

type Server interface {
	Start(ctx context.Context)
	Shutdown() error
}

type server struct {
	e  *echo.Echo
	db *sql.DB
}

type Conn struct {
}

func dbOpen(datasourceName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", datasourceName)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Duration(100) * time.Second)

	return db, nil
}

func NewServer(ctx context.Context) Server {
	e := echo.New()

	conn, err := dbOpen("root:pass@tcp(localhost:13306)/test")
	if err != nil {
		e.Logger.Fatal(err)
	}
	return &server{
		e:  e,
		db: conn,
	}
}

func (s *server) Start(ctx context.Context) {
	defer s.db.Close()

	app := &handlers.App{Db: s.db}
	opapp.RegisterHandlers(s.e, app)

	apiGroup := s.e.Group("")

	apiGroup.Use(middleware.Logger())
	apiGroup.Use(middleware.Recover())
	apiGroup.Use(middlewares.AuthHandler(s.db))

	go func() {
		log.Println("server start")
		s.e.Logger.Fatal(s.e.Start(":63342"))
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("stopping server...")
	if err := s.Shutdown(); err != nil {
		log.Fatalf("failed to stop server: %v", err)
	}
}

func (s *server) Shutdown() error {
	return s.db.Close()
}
