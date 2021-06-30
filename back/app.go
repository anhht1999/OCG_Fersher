package main

import (
	"github.com/anhht1999/back/config"
	"github.com/anhht1999/back/repo"
	"github.com/anhht1999/back/routes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		panic(err.Error())
	}
	repo.Connect(config.Config)
	// repo.InitMasterData()
	
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
	})
	app.Use(cors.New())
	app.Static("/public", "./public", fiber.Static{ //http://localhost:3000/public OR http://localhost:3000/public/dog.jpeg
		Compress:  true,
		ByteRange: true,
		Browse:    true,
		MaxAge:    3600,
	})

	routes.ConfigRouter(app)

	app.Listen(":3000")
}
