package controller

import (
	"fmt"
	"github/exmaximum/passwordmanager/src/helpers"
	"github/exmaximum/passwordmanager/src/models"
	"github/exmaximum/passwordmanager/src/services"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	services.UserService
}

type NewPassBody struct {
	Site   string
	Length int
}

func (u UserController) Register(c *fiber.Ctx) error {
	var user models.User
	c.BodyParser(&user)
	message := u.CreateUser(&user)
	if !message.Success {
		return c.Status(fiber.StatusBadRequest).JSON(message)
	}
	token, exp, err := helpers.CreateJwtToken(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(helpers.NewMessage(err.Error(), false))
	}
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Unix(exp, 0),
		HTTPOnly: true,
	})
	return c.Status(fiber.StatusCreated).JSON(message)
}

func (u UserController) Login(c *fiber.Ctx) error {
	var user models.User
	fmt.Println(c.IP())

	c.BodyParser(&user)
	message := u.SignUser(&user)
	if !message.Success {
		return c.Status(fiber.StatusBadRequest).JSON(message)
	}
	token, exp, err := helpers.CreateJwtToken(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Unix(exp, 0),
		HTTPOnly: true,
	})
	return c.Status(fiber.StatusOK).JSON(message)
}

func (u UserController) NewPassword(c *fiber.Ctx) error {
	// body := make(map[string]string, 1)
	body := NewPassBody{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewMessage(err.Error(), false))
	}
	if len(body.Site) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewMessage("Site is required", false))
	}

	if body.Length == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewMessage("Length is required", false))
	}
	userId := ""
	c.Request().Header.VisitAll(func(key, value []byte) {
		if string(key) == "Userid" {
			userId = string(value)
		}
	})
	password := u.AddPassword(userId, body.Site, body.Length)
	if !password.Success {
		return c.Status(fiber.StatusInternalServerError).JSON(password)
	}
	return c.Status(fiber.StatusOK).JSON(password)
}

func (u UserController) ChangeSitePassword(c *fiber.Ctx) error {
	body := NewPassBody{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewMessage(err.Error(), false))
	}
	if len(body.Site) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewMessage("Site is required", false))
	}
	if body.Length == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewMessage("Length is required", false))
	}
	userId := ""
	c.Request().Header.VisitAll(func(key, value []byte) {
		if string(key) == "Userid" {
			userId = string(value)
		}
	})
	password := u.UpdateSitePassword(userId, body.Site, body.Length)
	if !password.Success {
		return c.Status(fiber.StatusInternalServerError).JSON(password)
	}
	return c.Status(fiber.StatusOK).JSON(password)

}

func (u UserController) GetPassword(c *fiber.Ctx) error {
	queryParams := c.Query("site")
	if queryParams == "" {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewMessage("Site is required", false))
	}
	var userId string
	c.Request().Header.VisitAll(func(key, value []byte) {
		if string(key) == "Userid" {
			userId = string(value)
		}
	})
	password := u.GetUserPassword(queryParams, userId)
	if !password.Success {
		return c.Status(fiber.StatusInternalServerError).JSON(password)
	}
	return c.Status(fiber.StatusOK).JSON(password)
}

func (u UserController) GetUser(c *fiber.Ctx) error {

	var userId string
	c.Request().Header.VisitAll(func(key, value []byte) {
		if string(key) == "Userid" {
			userId = string(value)
		}
	})
	user := u.GetUserInfo(userId)
	if !user.Success {
		return c.Status(fiber.StatusInternalServerError).JSON(user)
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func (u UserController) Logout(c *fiber.Ctx) error {
	token := c.Cookies("token")
	if token == "" {
		return c.Status(fiber.StatusForbidden).JSON(helpers.NewMessage("You must have logged in", false))
	}
	c.Cookie(&fiber.Cookie{Name: "token", Value: "", Expires: time.Now().Add(-time.Hour), HTTPOnly: true})
	return c.Status(fiber.StatusOK).JSON(helpers.NewMessage("Logged out", true))
}
