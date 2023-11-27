package service

import (
	"bell-service/internal/model"
	"bell-service/internal/repository"
	"context"
	"github.com/pkg/errors"
	"strconv"
)

type rudService struct {
	db repository.MethodDatabase
}

func (r *rudService) GetByID(ctx context.Context, id string) (*model.BellInfo, error) {
	idi, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.Wrap(err, "GetByID: id is can't convert to int")
	}
	return r.db.GetBell(ctx, idi)
}

func (r *rudService) GetAll(ctx context.Context) ([]model.BellInfo, error) {
	return r.db.GetBells(ctx)
}

func (r *rudService) Update(ctx context.Context, bell *model.BellInfo) error {
	return r.db.UpdateBell(ctx, bell)
}

func (r *rudService) Delete(ctx context.Context, id string) error {
	idi, err := strconv.Atoi(id)
	if err != nil {
		return errors.Wrap(err, "DeleteBell: id is can't convert to int")
	}
	return r.db.DeleteBell(ctx, idi)
}

func newRUDService(db repository.MethodDatabase) *rudService {
	return &rudService{db: db}
}

var _ RUD = (*rudService)(nil)
