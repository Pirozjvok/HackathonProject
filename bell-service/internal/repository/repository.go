package repository

import (
	"bell-service/internal/model"
	"bell-service/pkg/database/postgres"
	"context"
)

type MethodDatabase interface {
	SetBell(ctx context.Context, bell *model.BellInfo) error
	UpdateBell(ctx context.Context, bell *model.BellInfo) error
	GetBell(ctx context.Context, id int) (*model.BellInfo, error)
	GetBells(ctx context.Context) ([]model.BellInfo, error)
	DeleteBell(ctx context.Context, id int) error
}

type Repositories struct {
	MDB MethodDatabase
}

func NewRepositories(db *postgres.Postgres) *Repositories {
	return &Repositories{
		MDB: newBellRepository(db),
	}
}
