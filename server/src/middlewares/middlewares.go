package middlewares

import (
	"github/exmaximum/passwordmanager/src/helpers"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/utils"

	"github.com/gofiber/helmet/v2"
)

func FiberMiddleware(app *fiber.App) {
	corsConfig := cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000, http://127.0.0.1:3000, http://192.168.1.23:3000",
		AllowCredentials: true,
	})

	app.Use(corsConfig, helmet.New(), limiterMiddleware(), encryptCookieMiddleware(), csrfMiddleware()) //  csrfMiddleware(),

}

func limiterMiddleware() func(*fiber.Ctx) error {
	limiterConfig := limiter.Config{
		Max:               100,
		Expiration:        15 * time.Minute,
		LimiterMiddleware: limiter.SlidingWindow{},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(helpers.NewMessage("Too many request", false))
		},
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		SkipSuccessfulRequests: false,
		SkipFailedRequests:     false,
	}
	return limiter.New(limiterConfig)
}

func csrfMiddleware() func(*fiber.Ctx) error {
	csrfConfig := csrf.Config{
		KeyLookup:      "header:X-Csrf-Token",
		CookieName:     "csrf_",
		CookieSameSite: "Strict",
		Expiration:     1 * time.Hour,
		KeyGenerator:   utils.UUID,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusForbidden).JSON(helpers.NewMessage("Forbidden", false))

		},
	}

	// Use middlewares for each route
	return csrf.New(csrfConfig)
}

func encryptCookieMiddleware() func(*fiber.Ctx) error {
	encryptCookieConfig := encryptcookie.Config{
		Key: os.Getenv("COOKIE_TOKEN"),
	}

	return encryptcookie.New(encryptCookieConfig)
}
