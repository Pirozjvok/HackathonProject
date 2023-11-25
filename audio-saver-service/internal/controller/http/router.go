package http

import (
	"audio-saver-service/internal/service"
	"github.com/gorilla/mux"
)

func NewRouter(rout *mux.Router, services *service.Services) {
	newAudioRoutes(services, rout)
}
