package service

import "github.com/gofiber/fiber/v2"

func Register(app *fiber.App, h *Handler) {
	app.Get("/queries", h.Get)
}
