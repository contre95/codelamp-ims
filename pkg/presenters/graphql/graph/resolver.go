package graph

import (
	"codelamp-ims/pkg/domain/clients"
	"codelamp-ims/pkg/domain/health"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	clientService clients.Service
	healthService health.Service
}
