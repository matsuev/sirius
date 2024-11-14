package handlers

import (
	"sirius-metrics/internal/metrics"

	"github.com/gofiber/fiber/v2"
)

// HomeHandler ...
func HomeHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		metrics.ObserveRequest(ctx.Path())

		return ctx.SendStatus(fiber.StatusOK)
	}
}

// HelpHandler ...
func HelpHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		metrics.ObserveRequest(ctx.Path())

		return ctx.SendString("help")
	}
}
