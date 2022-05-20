package handler

import "github.com/gofiber/fiber/v2"

type HealthHandler struct{}

func (hh *HealthHandler) Register(router fiber.Router) {
	router.Get("/", hh.getHealth)
}

func (hh *HealthHandler) getHealth(ctx *fiber.Ctx) error {
	return ctx.SendStatus(204)
}
