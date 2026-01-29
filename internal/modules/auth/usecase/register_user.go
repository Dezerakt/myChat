package usecase

import (
	"context"
	"fmt"
	"myChat/internal/modules/auth/domain"
	"myChat/internal/modules/auth/infrastructure/otp"
	"myChat/pkg/logs"
	"time"
)

func (obj *Usecase) RegisterUser(ctx context.Context, req domain.LogUpRequest) error {
	logger := logs.FromContext(ctx)

	key, err := obj.otp.GetCode(ctx, otp.GetCodeParams{
		Issuer:      obj.cfg.Auth.Issuer,
		AccountName: req.AccountName,
	})
	if err != nil {
		return fmt.Errorf("totp.GenerateCode: %w", err)
	}

	return nil
}
