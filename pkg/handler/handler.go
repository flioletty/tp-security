package handler

import (
	"net/http"
	"proxy/pkg/service"

	"github.com/gorilla/mux"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() http.Handler {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	apiRouter := api.PathPrefix("/v1").Subrouter()

	apiRouter.HandleFunc("/requests", h.getRequests).Methods("GET")
	apiRouter.HandleFunc("/request/{id}", h.getRequestById).Methods("GET")
	apiRouter.HandleFunc("/repeat/{id}", h.repeatRequestById).Methods("GET")
	apiRouter.HandleFunc("/scan/{id}", h.scanRequestById).Methods("GET")

	return r
}
