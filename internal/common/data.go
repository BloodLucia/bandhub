package common

import (
	"fmt"
	"log"

	"github.com/google/wire"
	"xorm.io/xorm"
)

type Data struct {
	DB *xorm.Engine
}

func NewData() (*Data, func(), error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&collation=utf8mb4_unicode_ci",
		"root",
		"root",
		"127.0.0.1",
		3306,
		"bandhub_dev",
	)
	db, err := xorm.NewEngine("mysql", dsn)

	if err != nil {
		return nil,nil, err
	}

	data := &Data{
		DB: db,
	}

	return data,func() {
		if err := db.Close();err!=nil {
			log.Panicf("Falied to close database: %s", err)
		}
	}, nil
}

var DataProvider = wire.NewSet(NewData)