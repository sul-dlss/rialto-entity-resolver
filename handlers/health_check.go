package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/sul-dlss/rialto-entity-resolver/generated/models"
	"github.com/sul-dlss/rialto-entity-resolver/generated/restapi/operations"
)

// NewHealthCheck will return the service health
func NewHealthCheck() operations.HealthCheckHandler {
	return &healthCheck{}
}

// healthCheck handles a request for the health of the service
type healthCheck struct{}

// Handle the retrieve resource request
func (d *healthCheck) Handle(params operations.HealthCheckParams) middleware.Responder {
	return operations.NewHealthCheckOK().WithPayload(&models.HealthCheckResponse{Status: "OK"})
}
