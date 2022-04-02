package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/peltastic/golang-jwt/controllers"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/login", controllers.Login)
	app.Get("/user", controllers.User)
}