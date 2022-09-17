package controllers

import (
	"net/http"
	"server/internal/app/services"

	"go.uber.org/zap"
)

type CoreDumpsRestController struct {
	service services.CoreDumpsService
	logger  *zap.Logger
}

func NewCoreDumpsRestController(s services.CoreDumpsService, l *zap.Logger) *CoreDumpsRestController {
	return &CoreDumpsRestController{
		service: s,
		logger:  l,
	}
}

func (c *CoreDumpsRestController) GetCoreDumps() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (c *CoreDumpsRestController) AddCoreDump() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (c *CoreDumpsRestController) DeleteCoreDump() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
