package main

import (
	"context"
	"log"
	"server/services"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func main() {

	//Create mux and setup to grpc
	ctx := context.TODO()
	mux := runtime.NewServeMux()
	//opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := services.RegisterCalculatorHandlerServer(ctx, mux, services.NewCalculationServer()); err != nil {
		log.Fatalf("failed to start HTTP gateway: %v\n", err)
	}

	//Create fiber app
	app := fiber.New()

	app.Get("/my-api", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	//Use go fiber adapter to make fiber and mux work together.
	app.Post("/my-api/grpc/*", adaptor.HTTPHandlerFunc(mux.ServeHTTP))

	app.Listen(":3000")
}
