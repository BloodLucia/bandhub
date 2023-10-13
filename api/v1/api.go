package v1

import (
	"github.com/google/wire"
	"github.com/kalougata/bandhub/internal/controller"
)

type APIRouter struct {
	userController controller.UserController
}

func NewAPIRouter(
	userController controller.UserController,
) *APIRouter {
	return &APIRouter{
		userController,
	}
}

var APIRouterProvider = wire.NewSet(NewAPIRouter)
