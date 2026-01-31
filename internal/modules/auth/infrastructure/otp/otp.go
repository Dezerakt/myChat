package otp

import "context"

type Otp interface {
	GetCode(ctx context.Context, params GetCodeParams) (string, error)
	GetQR(ctx context.Context, params GetQrParams) (string, error)
}

type GetQrParams struct {
}

type GetCodeParams struct {
	Issuer      string
	AccountName string
	Period      uint
}
