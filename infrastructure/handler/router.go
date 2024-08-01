package handler

import (
	"github.com/MikelSot/interseguro-challenge-qr/infrastructure/handler/qr"
	"github.com/MikelSot/interseguro-challenge-qr/model"
)

func InitRoutes(spec model.RouterSpecification) {
	// Q
	qr.NewRouter(spec)
}
