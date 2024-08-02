package qr

import (
	"github.com/gofiber/fiber/v2"

	"github.com/MikelSot/interseguro-challenge-qr/domain/qr"
	"github.com/MikelSot/interseguro-challenge-qr/model"
)

const (
	_privateRoutePrefix = "/qr/api/v1/factorize"
)

func NewRouter(spec model.RouterSpecification) {
	handler := buildHandler()

	privateRoutes(spec.App, handler)

}

func buildHandler() handler {
	useCase := qr.New()

	return newHandler(useCase)
}

func privateRoutes(app *fiber.App, handler handler, middlewares ...fiber.Handler) {
	api := app.Group(_privateRoutePrefix, middlewares...)

	api.Post("", handler.FactorizeQR)
}
