package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "hello!")
	})

	if err := r.Run(":3000"); err != nil {
		log.Panicf("failed to start http server: %s", err)
	}
}
