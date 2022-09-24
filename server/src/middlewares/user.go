package middlewares

import (
	"github/exmaximum/passwordmanager/src/helpers"
	"github/exmaximum/passwordmanager/src/models"
	"net/mail"

	"github.com/gofiber/fiber/v2"
)

func RegisterValidate(c *fiber.Ctx) error {
	var userInfos models.User

	if err := c.BodyParser(&userInfos); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	if len(userInfos.Email) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewMessage("Email is required", false))
	}

	if len(userInfos.Password) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewMessage("Password is required", false))
	}

	if len(userInfos.Username) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewMessage("Username is required", false))
	}

	_, err := mail.ParseAddress(userInfos.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewMessage("Email is not valid", false))
	}

	return c.Next()
}

func LoginValidate(c *fiber.Ctx) error {
	var userInfos models.User
	if err := c.BodyParser(&userInfos); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	if len(userInfos.Email) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewMessage("Email is required", false))
	}

	if len(userInfos.Password) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewMessage("Password is required", false))
	}

	_, err := mail.ParseAddress(userInfos.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewMessage("Email is not valid", false))
	}

	return c.Next()
}

func ValidateLogin(c *fiber.Ctx) error {
	token, err := helpers.VerifyToken(c.Cookies("token"))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(helpers.NewMessage(err.Error(), false))
	}
	c.Request().Header.Add("Userid", token.Id)
	return c.Next()
}
