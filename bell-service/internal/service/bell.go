package service

import (
	"bell-service/internal/model"
	"bell-service/internal/repository"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

type bellService struct {
	urlAPI     string
	db         repository.MethodDatabase
	serviceSHD string
}

func (b *bellService) Request(ctx context.Context) error {

	for i := 0; i < 6; i++ {
		resp, err := http.Get(b.urlAPI)
		if err != nil {
			return errors.Wrap(err, "Request: fail to GET request to API server")
		}

		bellInfo := &model.BellInfo{}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return errors.Wrap(err, "Request: fail to read response body")
		}

		err = json.Unmarshal(body, &bellInfo)
		if err != nil {
			return errors.Wrap(err, "Request: fail to json unmarshal BellInfo")
		}

		err = b.db.SetBell(ctx, bellInfo)
		if err != nil {
			return err
		}
	}

	return nil
}

func newBellService(urlAPI string, db repository.MethodDatabase, serviceSHD string) *bellService {
	return &bellService{urlAPI: urlAPI, db: db, serviceSHD: serviceSHD}
}

var _ APIRequester = (*bellService)(nil)
