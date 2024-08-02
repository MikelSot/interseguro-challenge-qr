package model

type Challenge struct {
	Factorize QR                `json:"factorize"`
	Statistic StatisticResponse `json:"statistic"`
}
