package usecase

import (
	"myChat/config"
	"myChat/internal/modules/auth/infrastructure/otp"
)

type Usecase struct {
	cfg *config.Config

	otp otp.Otp
}

func New(cfg *config.Config) *Usecase {
	return &Usecase{
		cfg: cfg,

		otp: otp.NewTotp(),
	}
}
