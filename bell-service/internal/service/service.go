package service

import (
	"bell-service/internal/model"
	"bell-service/internal/repository"
	"context"
)

// APIRequester is interface, which to create request in API server
type APIRequester interface {
	Request(ctx context.Context) error
}

type RUD interface {
	GetByID(ctx context.Context, id string) (*model.BellInfo, error)
	GetAll(ctx context.Context) ([]model.BellInfo, error)
	Update(ctx context.Context, bell *model.BellInfo) error
	Delete(ctx context.Context, id string) error
}

type Dependencies struct {
	UrlAPI     string
	DB         *repository.Repositories
	ServiceSHD string
}

type Services struct {
	APIRequester
	RUD
}

func NewServices(deps *Dependencies) *Services {
	return &Services{
		APIRequester: newBellService(deps.UrlAPI, deps.DB.MDB, deps.ServiceSHD),
		RUD:          newRUDService(deps.DB.MDB),
	}
}
