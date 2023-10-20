package adapter

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/marianozunino/go-scrap-challenge/internal/adapter/controller"
	"github.com/marianozunino/go-scrap-challenge/internal/port"
)

var _ port.App = &app{}

type app struct {
	controller *controller.Controller
	srv        *fiber.App
}

func NewApp() port.App {

	srv := fiber.New()

	service := NewService()
	controller := controller.NewAuthController(service)
	srv.Use(controller.IsAuthenticated)
	srv.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        60,
		Expiration: 60 * time.Second,
	}))

	a := app{
		srv:        srv,
		controller: controller,
	}

	a.registerRoutes()

	return &a
}

func (app *app) registerRoutes() {
	app.srv.Get("/", app.controller.Index)
	app.srv.Post("/", app.controller.Login)
	app.srv.Get("/orders", app.controller.Orders)
	app.srv.Get("/orders/:id/details", app.controller.OrderDetails)
	app.srv.Get("/logout", app.controller.Logout)
}

// Run implements port.App.
func (app *app) Run() {
	app.srv.Listen(":8080")
}
