package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ronaldalds/res/internal/services"
)

type Controller struct {
	Service *services.Service
}

func NewController() *Controller {
	return &Controller{
		Service: services.NewService(),
	}
}

func (con *Controller) HealthHandler(c *fiber.Ctx) error {
	return c.JSON(con.Service.Health())
}
