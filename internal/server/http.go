package server

import "github.com/gin-gonic/gin"

func NewServerHTTP() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "hello")
	})

	return r
}
