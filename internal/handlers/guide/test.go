package guide

import (
	"native/crus/pkg/helper"

	"github.com/gofiber/fiber/v2"
)

func Test(c *fiber.Ctx) error {
	return helper.ResponseJson(c, 200, "Sukses", nil)
}
