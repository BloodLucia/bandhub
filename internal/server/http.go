package server

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/kalougata/bandhub/api/v1"
)

func NewServerHTTP(
	apiRouter *v1.APIRouter,
) *gin.Engine {
	r := gin.Default()
	rootGroup := r.Group("")

	// 无需鉴权的路由
	unAuthGroup := rootGroup.Group("/api/v1")
	apiRouter.RegisterGuestAPIRouter(unAuthGroup)

	return r
}
