package model

type Statistic struct {
	MaxValue   float64 `json:"max_value"`
	MinValue   float64 `json:"min_value"`
	Average    float64 `json:"average"`
	TotalSum   float64 `json:"total_sum"`
	IsDiagonal bool    `json:"is_diagonal"`
}

type StatisticResponse struct {
	StatisticQ Statistic `json:"statistic_q"`
	StatisticR Statistic `json:"statistic_r"`
}
