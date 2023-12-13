package main

import (
	"categori/initializers"
	"categori/router"
	"github.com/gofiber/fiber/v2"
	"log"
)

func init() {
	initializers.InitMongoDB()
	initializers.ConnectDb()
}

func main() {
	app := fiber.New()

	router.CorsConfig(app)

	router.ConfigureAPIRoutes(app)

	log.Fatal(app.Listen(":8000"))
}
