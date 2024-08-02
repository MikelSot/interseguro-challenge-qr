package handler

import (
	"github.com/MikelSot/interseguro-challenge-qr/infrastructure/handler/challenge"
	"github.com/MikelSot/interseguro-challenge-qr/infrastructure/handler/qr"
	"github.com/MikelSot/interseguro-challenge-qr/model"
)

func InitRoutes(spec model.RouterSpecification) {
	// C
	challenge.NewRouter(spec)

	// Q
	qr.NewRouter(spec)
}
