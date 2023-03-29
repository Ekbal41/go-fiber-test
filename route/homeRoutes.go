package route

import (
	"github.com/gofiber/fiber/v2"
	"github.cpm/ekbal41/fiber-test-1/controller"
)

func HomeRoutes(route fiber.Router) {
	route.Get("/", controller.HomePage)
	route.Get("/about", controller.AboutPage)
}