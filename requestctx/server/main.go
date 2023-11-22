package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func main() {
	go fiberServer()
	ginServer()
}

func fiberServer() {
	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		select {
		case <-ctx.Done():
			fmt.Println("fiber success")
		case <-c.Context().Done():
			fmt.Println("fiber connection lost")
		}

		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":9090")
}

func ginServer() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		select {
		case <-ctx.Done():
			fmt.Println("gin success")
		case <-c.Request.Context().Done():
			fmt.Println("gin connection lost")
		}
	})

	r.Run()
}
