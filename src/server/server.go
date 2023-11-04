package server

import (
	"onepixel_backend/src/controllers"
	"gorm.io/gorm"
	"onepixel_backend/src/routes/api"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateApp(dbConn *gorm.DB) *fiber.App {
	app := fiber.New()

	usersController := controllers.NewUsersController(dbConn)
	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	apiV1 := app.Group("/api/v1")

	apiV1.Route("/users", func(router fiber.Router) {
        api.UsersRoute(router, usersController)
    })

	apiV1.Route("/urls", api.UrlsRoute)

	return app
}
