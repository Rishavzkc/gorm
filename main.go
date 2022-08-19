package main

import (
	"mysqlgin/config"
	"mysqlgin/database"
	"mysqlgin/routes"
)

func main() {
	config.Init()

	database.StartDatabase()
	s := routes.NewServer()

	s.Run()
}
