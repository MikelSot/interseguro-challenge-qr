package domain

type UseCase interface {
	GenerateQR() (string, error)
}
