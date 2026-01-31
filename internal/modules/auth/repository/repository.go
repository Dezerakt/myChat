package repository

import "context"

type Profile interface {
	SaveNewProfile(ctx context.Context) error
}
