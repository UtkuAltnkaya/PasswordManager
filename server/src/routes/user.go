package routes

import (
	"github/exmaximum/passwordmanager/src/controller"
	"github/exmaximum/passwordmanager/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

type UserRouter struct {
	controller.UserController
}

func (user UserRouter) UserRoute(app *fiber.App) {

	app.Post("/register", middlewares.RegisterValidate, user.Register)
	app.Post("/logout", user.Logout)
	app.Post("/login", middlewares.LoginValidate, user.Login)
	app.Post("/generatePassword", middlewares.ValidateLogin, user.NewPassword)
	app.Get("/getUser", middlewares.ValidateLogin, user.GetUser)
	app.Get("/getPassword", middlewares.ValidateLogin, user.GetPassword)
}
