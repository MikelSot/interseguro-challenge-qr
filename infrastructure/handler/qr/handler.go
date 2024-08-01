package qr

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"

	"github.com/MikelSot/interseguro-challenge-qr/domain/qr"
	"github.com/MikelSot/interseguro-challenge-qr/model"
)

type handler struct {
	useCase qr.UseCase
}

func newHandler(uc qr.UseCase) handler {
	return handler{uc}
}

func (h handler) FactorizeQR(c *fiber.Ctx) error {
	payload := model.Matrix{}

	if err := c.BodyParser(&payload); err != nil {
		log.Warn("¡Uy! Error al leer el cuerpo de la solicitud", err.Error())

		return c.Status(fiber.StatusBadRequest).JSON(model.MessageResponse{
			Errors: model.Responses{
				{Code: model.BindFailed, Message: "Error al leer el cuerpo de la solicitud"},
			},
		})
	}

	factorizeQr, err := h.useCase.FactorizeQR(payload.Matrix)
	if err != nil {
		log.Warn("qr: usecase.FactorizeQR()", err.Error())

		customErr := model.NewError()
		if errors.As(err, &customErr) {
			return customErr
		}

		return c.Status(fiber.StatusInternalServerError).JSON(model.MessageResponse{
			Errors: model.Responses{
				{Code: model.Failure, Message: "Error al factorizar la matriz"},
			},
		})
	}

	return c.Status(http.StatusOK).JSON(model.MessageResponse{Data: factorizeQr})
}
