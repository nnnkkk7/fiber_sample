package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	now := time.Now()
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	fmt.Printf("経過: %vms\n", time.Since(now).Nanoseconds())
	app.Listen(":3000")
}
