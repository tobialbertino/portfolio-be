package main

import (
	"log"
	"os"
	appConfig "tobialbertino/portfolio-be/app"
	"tobialbertino/portfolio-be/exception"
	"tobialbertino/portfolio-be/pkg/config"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	var (
		cfg      *config.Config
		err      error
		validate *validator.Validate = validator.New()
	)

	// Load config
	cfg, err = config.LoadConfig()
	if err != nil {
		log.Printf("error loading config: %s", err)
	}

	// Add DB
	DB := appConfig.NewDB(cfg)
	defer DB.Close()

	// Use default logger
	file, err := os.OpenFile("./info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()
	log.SetOutput(file)

	// initiate framework
	app := fiber.New(fiber.Config{
		ErrorHandler: exception.CustomErrorHandler,
	})
	// Recover Panic
	app.Use(recover.New())
	// logger Middleware
	app.Use(logger.New(logger.Config{
		Output: file,
	}))

	// set modules & app Router
	appConfig.InitRouter(app, DB, validate)

	if cfg.Server.Port == "" {
		log.Println("Port tidak ditemukan")
	}
	app.Listen(":" + cfg.Server.Port)
}
