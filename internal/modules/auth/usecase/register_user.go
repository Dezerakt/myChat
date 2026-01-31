package usecase

import (
	"context"
	"myChat/internal/modules/auth/domain"
	"myChat/pkg/logs"
)

func (obj *Usecase) RegisterUser(ctx context.Context, req *domain.LogUpRequest) error {
	logger := logs.FromContext(ctx)

	logger.Debug().Interface("req", req).Msg("RegisterUser")

	err := obj.profileRepo.SaveNewProfile(ctx)
	if err != nil {
		logger.Debug().Err(err).Msg("RegisterUser")
		return err
	}

	return nil
}
