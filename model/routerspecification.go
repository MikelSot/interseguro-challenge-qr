package model

import (
	"github.com/gofiber/fiber/v2"
)

type RouterSpecification struct {
	App             *fiber.App
	Logger          Logger
	ConfigStatistic ConfigStatistic
}
