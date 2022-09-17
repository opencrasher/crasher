package controllers

import (
	"encoding/json"
	"net/http"
	"server/internal/app/services"

	"go.uber.org/zap"
)

type ApplicationsRestController struct {
	service services.ApplicationsService
	logger  *zap.Logger
}

func NewApplicationsRestController(s services.ApplicationsService, l *zap.Logger) *ApplicationsRestController {
	return &ApplicationsRestController{
		service: s,
		logger:  l,
	}
}

func (c *ApplicationsRestController) GetApplicationNames() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		appNames, err := c.service.GetApplicationNames()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			c.logger.Error("get application names failed", zap.Error(err))
			return
		}

		response, err := json.Marshal(appNames)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			c.logger.Error("response marshaling failed", zap.Error(err))
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(response)
	}
}
