package db

import (
	"github.com/dgraph-io/badger/v4"
	"log"
)

const pathName = "./db_table" //数据库持久化地址，默认"./db_table"
var DB *badger.DB

func InitDB() {
	var err error
	DB, err = badger.Open(badger.DefaultOptions(pathName).WithInMemory(false))
	if err != nil {
		log.Fatal(err)
	}
}
