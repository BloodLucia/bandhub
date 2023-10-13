package common

import (
	"github.com/google/wire"
	"github.com/kalougata/bandhub/internal/controller"
)

var ControllerProvider = wire.NewSet(controller.NewUserController)