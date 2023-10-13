//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	v1 "github.com/kalougata/bandhub/api/v1"
	"github.com/kalougata/bandhub/internal/common"
)

func NewApp() (*common.Server, func(), error) {
	panic(wire.Build(
		common.ControllerProvider,
		v1.APIRouterProvider,
		common.ServerProvider,
	))
}
