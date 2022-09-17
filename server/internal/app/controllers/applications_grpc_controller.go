package controllers

import (
	"server/internal/app/services"

	"go.uber.org/zap"
)

type ApplicationsGrpcController struct {
	service services.ApplicationsService
	logger  *zap.Logger
}

func NewApplicationsGrpcController(s services.ApplicationsService, l *zap.Logger) *ApplicationsGrpcController {
	return &ApplicationsGrpcController{
		service: s,
		logger:  l,
	}
}
