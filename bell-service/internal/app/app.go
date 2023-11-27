package app

import (
	"bell-service/config"
	"bell-service/internal/controller/http"
	"bell-service/internal/repository"
	"bell-service/internal/service"
	"bell-service/pkg/database/postgres"
	"bell-service/pkg/server"
	"context"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	ctx := context.Background()

	pg, err := postgres.New(ctx, cfg.PG.ConnURI)
	if err != nil {
		log.Println(err)
	}

	repositories := repository.NewRepositories(pg)

	deps := &service.Dependencies{UrlAPI: cfg.API.UrlAPI, DB: repositories}

	services := service.NewServices(deps)
	rout := mux.NewRouter()
	http.NewRouter(rout, services)
	srv := server.NewServer(rout, server.Port(cfg.HTTP.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Println("Run: " + s.String())
	case err := <-srv.Notify():
		log.Println(errors.Wrap(err, "Run: signal.Notify"))
	}

	err = srv.Shutdown()
	if err != nil {
		log.Println(errors.Wrap(err, "Run: server shutdown"))
	}
}
