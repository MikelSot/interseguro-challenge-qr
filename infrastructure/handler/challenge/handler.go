package challenge

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"

	"github.com/MikelSot/interseguro-challenge-qr/domain/challenge"
	"github.com/MikelSot/interseguro-challenge-qr/model"
)

type handler struct {
	useCase challenge.UseCase
}

func newHandler(cu challenge.UseCase) handler {
	return handler{cu}
}

func (h handler) Challenge(c *fiber.Ctx) error {
	payload := model.Matrix{}
	if err := c.BodyParser(&payload); err != nil {
		log.Warn("challenge: ¡Uy! Error al leer el cuerpo de la solicitud", err.Error())

		return c.Status(fiber.StatusBadRequest).JSON(model.MessageResponse{
			Errors: model.Responses{
				{Code: model.BindFailed, Message: "Error al leer el cuerpo de la solicitud"},
			},
		})
	}

	m, err := h.useCase.Challenge(payload.Matrix)
	if err != nil {
		log.Warn("challenge: useCase.Challenge(): ", err.Error())

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

	return c.Status(http.StatusOK).JSON(model.MessageResponse{Data: m})
}
