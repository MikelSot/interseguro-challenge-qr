package bootstrap

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"github.com/MikelSot/interseguro-challenge-qr/infrastructure/handler"
	"github.com/MikelSot/interseguro-challenge-qr/infrastructure/handler/response"
	"github.com/MikelSot/interseguro-challenge-qr/model"
)

func Run() {
	_ = godotenv.Load()

	app := newFiber(response.ErrorHandler)
	logger := newLogger(false)

	handler.InitRoutes(model.RouterSpecification{
		App:    app,
		Logger: logger,
	})

	log.Fatal(app.Listen(getPort()))
}
