package challenge

import (
	"github.com/gofiber/fiber/v2"

	"github.com/MikelSot/interseguro-challenge-qr/domain/challenge"
	"github.com/MikelSot/interseguro-challenge-qr/domain/qr"
	"github.com/MikelSot/interseguro-challenge-qr/infrastructure/statistic"
	"github.com/MikelSot/interseguro-challenge-qr/model"
)

const (
	_privateRoutePrefix = "/qr/api/v1/challenge"
)

func NewRouter(spec model.RouterSpecification) {
	handler := buildHandler(spec)

	privateRoutes(spec.App, handler)

}

func buildHandler(spec model.RouterSpecification) handler {
	factorizeUseCase := qr.New()

	statisticService := statistic.New(spec.ConfigStatistic)

	useCase := challenge.New(factorizeUseCase, statisticService)

	return newHandler(useCase)
}

func privateRoutes(app *fiber.App, handler handler, middlewares ...fiber.Handler) {
	api := app.Group(_privateRoutePrefix, middlewares...)

	api.Post("", handler.Challenge)
}
