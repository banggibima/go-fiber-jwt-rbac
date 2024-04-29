package fiber

import (
	"os"

	"github.com/banggibima/go-fiber-jwt-rbac/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func New(config *config.Config) (*fiber.App, error) {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "fiber",
		AppName:       config.App.Name,
	})

	app.Use(logger.New(logger.Config{
		Output:        os.Stdout,
		DisableColors: false,
	}))

	return app, nil
}
