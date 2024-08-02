package challenge

import "github.com/MikelSot/interseguro-challenge-qr/model"

type Challenge struct {
	factorize QRUseCase
	statistic StatisticUseCase
}

func New(factorize QRUseCase, statistic StatisticUseCase) Challenge {
	return Challenge{
		factorize: factorize,
		statistic: statistic,
	}
}

func (c Challenge) Challenge(matrix [][]float64) (model.Challenge, error) {
	qr, err := c.factorize.FactorizeQR(matrix)
	if err != nil {
		return model.Challenge{}, err
	}

	statistic, err := c.statistic.GetStatistic(qr)
	if err != nil {
		return model.Challenge{}, err
	}

	return model.Challenge{
		Factorize: qr,
		Statistic: statistic,
	}, nil
}
