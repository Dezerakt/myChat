package otp

import (
	"context"
	"myChat/pkg/logs"

	"github.com/pquerna/otp/totp"
)

type totpOtp struct {
}

func NewTotp() Otp {
	return &totpOtp{}
}

func (obj *totpOtp) GetQR(ctx context.Context, params GetQrParams) (string, error) {
	return "", nil
}

func (obj *totpOtp) GetCode(ctx context.Context, params GetCodeParams) (string, error) {
	logger := logs.FromContext(ctx)

	logger.Debug().Interface("Request to get new OTP code for accountName", params.AccountName).Msg("GetCode")
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      params.Issuer,
		AccountName: params.AccountName,
		Period:      params.Period,
		SecretSize:  0,
		Secret:      nil,
		Digits:      0,
		Algorithm:   0,
		Rand:        nil,
	})
	if err != nil {
		logger.Error().Err(err).Msg("Failed to generate OTP code")
		return "", err
	}

	return key.Secret(), nil
}
