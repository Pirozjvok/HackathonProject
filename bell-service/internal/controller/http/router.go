package http

import (
	"bell-service/internal/service"
	"github.com/gorilla/mux"
)

func NewRouter(rout *mux.Router, services *service.Services) {
	newBellRoutes(services, rout)
}
