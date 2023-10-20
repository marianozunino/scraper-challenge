package controller

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/marianozunino/go-scrap-challenge/internal/core/dto"
	"github.com/marianozunino/go-scrap-challenge/internal/port"
	"github.com/marianozunino/go-scrap-challenge/view"
	"github.com/marianozunino/go-scrap-challenge/view/layout"
)

type Controller struct {
	service port.Service
	store   *session.Store
}

func sendHTML(c *fiber.Ctx, tmpl templ.Component) error {
	c.Type("html")
	return layout.Layout(tmpl).Render(c.Context(), c)
}

func NewAuthController(service port.Service) *Controller {
	cfg := session.Config{
		Expiration:     1 * time.Hour,
		CookieHTTPOnly: true,
	}
	store := session.New(cfg)
	return &Controller{
		service,
		store,
	}
}

func (cnt *Controller) Index(c *fiber.Ctx) error {
	index := view.Index("")
	return sendHTML(c, index)

}

func (cnt *Controller) Login(c *fiber.Ctx) error {
	p := new(dto.Login)
	if err := c.BodyParser(p); err != nil {
		return err
	}

	if cnt.service.Login(p.Username, p.Password) {
		session, err := cnt.store.Get(c)
		if err != nil {
			return err
		}

		session.Set("username", p.Username)
		session.Save()

		return c.Redirect("/orders")
	}

	return sendHTML(c, view.Index("Usuario o contraseña incorrectos"))
}

func (cnt *Controller) Orders(c *fiber.Ctx) error {
	pageStr := c.Query("page", "1")
	pageInt, err := strconv.Atoi(pageStr)

	if err != nil {
		return err
	}

	pageSize := 10

	orders, total := cnt.service.GetOrders(pageInt, pageSize)

	session, err := cnt.store.Get(c)
	if err != nil {
		return err
	}

	c.Type("html")

	more := pageInt*pageSize < int(total)

	index := view.Orders(orders, session.Get("username").(string), pageInt, more)
	return sendHTML(c, index)
}

func (cnt *Controller) OrderDetails(c *fiber.Ctx) error {
	orderIDStr := c.Params("id")
	orderIDInt, err := strconv.Atoi(orderIDStr)
	if err != nil {
		return err
	}

	bytes, err := cnt.service.GetOrderCSV(orderIDInt)
	if err != nil {
		return sendHTML(c, view.NotFound())
	}

	c.Set("Content-Type", "text/csv")
	c.Set("Content-Disposition", "attachment; filename=order.csv")
	c.Set("Content-Length", strconv.Itoa(len(bytes)))

	time.Sleep(time.Duration(1+rand.Intn(5)) * time.Second)

	return c.Send(bytes)
}

func (cnt *Controller) Logout(c *fiber.Ctx) error {
	session, err := cnt.store.Get(c)
	if err != nil {
		return err
	}
	session.Destroy()
	return c.Redirect("/")
}

func (cnt *Controller) IsAuthenticated(c *fiber.Ctx) error {
	session, err := cnt.store.Get(c)
	if err != nil {
		return err
	}
	if session.Get("username") != nil {
		if c.Path() != "/" {
			return c.Next()
		}
		return c.Redirect("/orders")
	} else {
		if c.Path() == "/" {
			return c.Next()
		}
		return c.Redirect("/")
	}
}
