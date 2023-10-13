package common

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/kalougata/bandhub/internal/server"
)

type Server struct {
	ServerHTTP *gin.Engine
}

func NewServer(
	ServerHTTP *gin.Engine,
) *Server {
	return &Server{ServerHTTP}
}

var ServerProvider = wire.NewSet(
	NewServer,
	server.NewServerHTTP,
)
