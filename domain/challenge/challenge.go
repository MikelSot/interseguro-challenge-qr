package challenge

import "github.com/MikelSot/interseguro-challenge-qr/model"

type UseCase interface {
	Challenge(matrix [][]float64) (model.Challenge, error)
}

type QRUseCase interface {
	FactorizeQR(matrix [][]float64) (model.QR, error)
}

type StatisticUseCase interface {
	GetStatistic(QR model.QR) (model.StatisticResponse, error)
}
