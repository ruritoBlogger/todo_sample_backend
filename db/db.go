package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ruritoBlogger/todo_sample_backend/entity"
	"gopkg.in/gorp.v2"
)

var (
	dbmap *gorp.DbMap
)

func Initialize() {
	// database/sql を用いて読み込む
	db, err := sql.Open("sqlite3", "/home/rurito/mysketch/todo_sample_backend/db/post_db.bin")

	if err != nil {
		log.Fatalln("Failed to open sql file", err)
	}

	// gorpを用いて取り込む
	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	// migration
	dbmap.AddTableWithName(entity.Post{}, "posts").SetKeys(true, "Id")
	err = dbmap.CreateTablesIfNotExists()

	if err != nil {
		log.Fatalln("Failed to create table", err)
	}

	log.Println("Successed to create db")
}

func GetDB() *gorp.DbMap {
	return dbmap
}

func Close() {
	dbmap.Db.Close()
}
