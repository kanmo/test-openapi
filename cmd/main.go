package main

import (
	"context"
	"log"
	"test-openapi/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	svr := server.NewServer(ctx)
	svr.Start(ctx)

	log.Println("server stopped")
}
