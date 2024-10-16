package http

import (
	"url-shortening-api/internal/link"
	"url-shortening-api/internal/user"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	UserController *user.UserController
	LinkController *link.LinkController
}

func (route *RouteConfig) NewRouter() *fiber.App {
	app := fiber.New()

	app.Post("/v1/users", route.UserController.Register)
	app.Post("/v1/users/_login", route.UserController.Login)
	app.Post("/v1/users/_logout", route.UserController.Logout)
	app.Put("/v1/users/_current", route.UserController.Update)
	app.Delete("/v1/users/_current", route.UserController.Delete)

	app.Get("/short/:id", route.LinkController.RedirectLink)
	app.Post("/v1/short", route.LinkController.Create)

	return app
}
