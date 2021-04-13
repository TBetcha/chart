package routes

import (
	"../controllers"
	"../middleware"
	"github.com/gofiber/fiber"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	//Private routes - use middleware

	app.Use(middleware.IsAuthenticated)

	app.Get("/api/user", controllers.User)
	app.Post("api/logout", controllers.Logout)
}
