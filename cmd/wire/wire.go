//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/kalougata/bandhub/internal/common"
)

func NewApp() (*common.Server, func(), error) {
	panic(wire.Build(
		common.ServerProvider,
	))
}
