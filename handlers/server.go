package handlers

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
	"regexp"
	"test-openapi/generated/models"
	opapp "test-openapi/generated/openapi/app"
	"test-openapi/repository"
)

var _ opapp.ServerInterface = (*App)(nil)

type App struct {
	Db *sql.DB
}

func (a App) CreateSpace(c echo.Context) error {
	log.Println("-------------- CreateSpace --------------")
	ctx := c.Request().Context()
	body := opapp.Space{}
	if err := c.Bind(&body); err != nil {
		return err
	}

	userName := c.Get("subject").(string)
	if userName != body.Owner {
		log.Printf("owner is invalid: %v", body.Owner)
		return echo.NewHTTPError(400, "owner is invalid")
	}

	repo := repository.NewSpaceRepository(a.Db)
	space := &models.Space{
		SpaceID: body.SpaceId,
		Name:    body.Name,
		Owner:   body.Owner,
	}

	err := repo.Create(ctx, space)
	if err != nil {
		log.Printf("failed to create space: %v", err)
		return echo.NewHTTPError(500, "failed to create space")
	}
	return nil

}

func (a App) CreateMessage(ctx echo.Context, spaceId int) error {
	//TODO implement me
	panic("implement me")
}

func (a App) ReadMessage(ctx echo.Context, spaceId int, messageId int) error {
	//TODO implement me
	panic("implement me")
}

func (a App) RegisterUser(c echo.Context) error {

	log.Println("-------------- RegisterUser --------------")
	ctx := c.Request().Context()
	body := opapp.User{}
	if err := c.Bind(&body); err != nil {
		return err
	}
	re := regexp.MustCompile(`^[a-zA-Z0-9_!@#$%^&*()]+$`)
	if !re.MatchString(body.Password) {
		log.Printf("password is invalid: %v", body.Password)
		return echo.NewHTTPError(400, "password is invalid")
	}

	hsPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("failed to hash password: %v", err)
		return echo.NewHTTPError(500, "failed to hash password")
	}

	repo := repository.NewUserRepository(a.Db)
	user := &models.User{
		ID:     body.Id,
		Name:   body.Name,
		PWHash: string(hsPassword),
	}

	err = repo.Create(ctx, user)
	if err != nil {
		log.Printf("failed to create user: %v", err)
		return echo.NewHTTPError(500, "failed to create user")
	}

	return nil

}

func (a App) DeleteUser(ctx echo.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (a App) GetUser(c echo.Context, id int) error {
	log.Println("-------------- GetUser --------------")

	ctx := c.Request().Context()
	repo := repository.NewUserRepository(a.Db)
	user, err := repo.FindByID(ctx, id)
	if err != nil {
		return echo.NewHTTPError(404, "user not found")
	}

	return c.JSON(200, opapp.User{
		Id:   user.ID,
		Name: user.Name,
	})
}
