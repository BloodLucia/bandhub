package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kalougata/bandhub/pkg/http"
	"xorm.io/xorm"
)

var DB *xorm.Engine

func connectDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&collation=utf8mb4_unicode_ci",
		"root",
		"root",
		"127.0.0.1",
		3306,
		"bandhub_dev",
	)
	db, err := xorm.NewEngine("mysql", dsn)

	if err != nil {
		log.Panicf("falied to connect database: %s", err)
	}

	log.Println("connected database")

	DB = db
}

func main() {
	r := gin.Default()

	connectDb()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "hello")
	})

	http.Run(r, ":8000")

	defer DB.Close()
}
