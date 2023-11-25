package app

import (
	"audio-saver-service/config"
	"audio-saver-service/internal/controller/http"
	"audio-saver-service/internal/service"
	"audio-saver-service/pkg/server"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {

	deps := &service.Dependencies{URL: cfg.API.Uri}

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

	err := srv.Shutdown()
	if err != nil {
		log.Println(errors.Wrap(err, "Run: server shutdown"))
	}
}
