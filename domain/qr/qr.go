package qr

import "github.com/MikelSot/interseguro-challenge-qr/model"

type UseCase interface {
	FactorizeQR(matrix [][]float64) (model.QR, error)
}
