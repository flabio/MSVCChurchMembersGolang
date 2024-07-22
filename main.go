package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"member_church.com/infrastructure/routers"
)

func main() {
	app := fiber.New()
	// Custom CORS configuration

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:8080", // Specify allowed origins
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Content-Type, Authorization",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true, // Allow credentials
	}))
	//app.Get("/swagger/*", swagger.HandlerDefault) // default
	app.Get("/swagger/*", swagger.HandlerDefault)
	routers.NewRouter(app)
	routers.NewMinisterialRouter(app)
	routers.NewTeamPescaRouter(app)
	routers.NewChurchRouter(app)
	routers.NewUserRouter(app)
	app.Listen(":8084")
}
