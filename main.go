package main

import (
	"github.com/ruritoBlogger/todo_sample_backend/db"

	"github.com/ruritoBlogger/todo_sample_backend/server"
)

func main() {
	db.Initialize()
	server.Start()
}
