package repository

import (
	"bell-service/internal/model"
	"bell-service/pkg/database/postgres"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

type bellRepository struct {
	db *postgres.Postgres
}

func (b *bellRepository) SetBell(ctx context.Context, bell *model.BellInfo) error {
	sql, args, err := b.db.Builder.Insert("bells").Columns("status", "callDate", "operatorFio", "clientPhone", "contactAudioURL", "statusBell").
		Values(bell.Status, bell.CallDate, bell.OperatorFio, bell.ClientPhone, bell.ContactAudio, "new").ToSql()
	if err != nil {
		return errors.Wrap(err, "SetBell: fail to create sql query")
	}

	_, err = b.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "SetBell: fail to insert bell")
	}

	return nil
}

func (b *bellRepository) UpdateBell(ctx context.Context, bell *model.BellInfo) error {
	sql, args, err := b.db.Builder.Update("bells").Set("statusBell", bell.StatusBell).Where(squirrel.Eq{"id": bell.ID}).ToSql()
	if err != nil {
		return errors.Wrap(err, "UpdateBell: fail to create sql query")
	}

	_, err = b.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "UpdateBell: fail to update bell")
	}
	return nil
}

func (b *bellRepository) GetBell(ctx context.Context, id int) (*model.BellInfo, error) {
	sql, args, err := b.db.Builder.Select("*").
		From("bells").Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "GetBell: fail to create sql query")
	}
	bell := &model.BellInfo{}

	row := b.db.Pool.QueryRow(ctx, sql, args...)
	err = row.Scan(&bell.ID, &bell.Status, &bell.CallDate, &bell.OperatorFio, &bell.ClientPhone, &bell.ContactAudio, &bell.StatusBell)
	if err != nil {
		return nil, errors.Wrap(err, "GetBell: fail to select")
	}

	return bell, nil
}

func (b *bellRepository) GetBells(ctx context.Context) ([]model.BellInfo, error) {
	sql, _, err := b.db.Builder.Select("*").
		From("bells").ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "GetBells: fail to create sql query")
	}

	rows, err := b.db.Pool.Query(ctx, sql)
	if err != nil {
		return nil, errors.Wrap(err, "GetBells: fail to find row")
	}
	defer rows.Close()

	entities := make([]model.BellInfo, 0)

	for rows.Next() {
		e := model.BellInfo{}

		err = rows.Scan(&e.ID, &e.Status, &e.CallDate, &e.OperatorFio, &e.ClientPhone, &e.ContactAudio, &e.StatusBell)
		if err != nil {
			return nil, errors.Wrap(err, "GetBells: fail to select")
		}

		entities = append(entities, e)
	}

	return entities, nil
}

func (b *bellRepository) DeleteBell(ctx context.Context, id int) error {
	sql, args, err := b.db.Builder.Delete("bells").Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return errors.Wrap(err, "DeleteBell: fail to create sql query")
	}

	_, err = b.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "DeleteBell: fail to remove bell")
	}
	return nil
}

func newBellRepository(db *postgres.Postgres) *bellRepository {
	return &bellRepository{db: db}
}
