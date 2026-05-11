package core

import "time"

const (
	DatabaseDir  = "database"
	DatabaseFile = "todos.db"
)

const (
	ServerAddress = "localhost:8082"
	ReadTimeout   = 5 * time.Second
	WriteTimeout  = 10 * time.Second
	IdleTimeout   = 60 * time.Second
)
