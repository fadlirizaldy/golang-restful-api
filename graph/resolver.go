package graph

import (
	"project_alterra/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	castService service.CastService
}
