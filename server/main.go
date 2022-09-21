package main

import (
	"github/exmaximum/passwordmanager/config"
	"github/exmaximum/passwordmanager/src/controller"
	"github/exmaximum/passwordmanager/src/middlewares"
	"github/exmaximum/passwordmanager/src/routes"
	"github/exmaximum/passwordmanager/src/services"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	config.LoadEnv()
	config.ConnectDb()
	app := fiber.New()

	file, err := os.OpenFile("./log/logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	// Set config for logger
	loggerConfig := logger.Config{
		Format: "[${time}] [${ip}]:${port} ${status} - ${method} ${latency} ${path}\n",
		Output: file, // add file to save output
	}

	// Use middlewares for each route
	app.Use(logger.New(loggerConfig))

	middlewares.FiberMiddleware(app)
	routes.UserRouter{UserController: controller.UserController{UserService: services.NewUserService(config.DataBase)}}.UserRoute(app)
	log.Fatal(app.Listen(os.Getenv("SERVER_PORT")))
}
