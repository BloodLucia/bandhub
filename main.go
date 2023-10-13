package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kalougata/bandhub/cmd/wire"
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

	servers, cleanup, err := wire.NewApp()

	if err != nil {
		log.Fatalln(err)
	}

	connectDb()

	http.Run(servers.ServerHTTP, ":8000")

	defer DB.Close()
	defer cleanup()
}
