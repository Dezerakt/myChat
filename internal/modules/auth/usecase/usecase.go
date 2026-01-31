package usecase

import (
	"myChat/config"
	"myChat/internal/modules/auth/infrastructure/otp"
	"myChat/internal/modules/auth/repository"
	pgRepo "myChat/internal/modules/auth/repository/postgres"
)

type Usecase struct {
	cfg *config.Config

	otp otp.Otp

	profileRepo repository.Profile
}

func New(cfg *config.Config) *Usecase {
	return &Usecase{
		cfg: cfg,

		otp: otp.NewTotp(),

		profileRepo: pgRepo.New(cfg),
	}
}
