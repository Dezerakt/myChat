package otp

import "context"

type Otp interface {
	GetCode(ctx context.Context, params GetCodeParams) (string, error)
}

type GetCodeParams struct {
	Issuer      string
	AccountName string
}
