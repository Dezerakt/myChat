package pgRepo

import (
	"context"
	"myChat/config"
	"myChat/internal/modules/auth/infrastructure/pg"

	"gorm.io/gorm"
)

type Profile struct {
	*gorm.DB
}

func New(cfg *config.Config) *Profile {
	connection, err := pg.NewConnection(&cfg.Postgres)
	if err != nil {
		return nil
	}

	return &Profile{
		DB: connection,
	}
}

func (obj *Profile) SaveNewProfile(ctx context.Context) error {
	return nil
}
