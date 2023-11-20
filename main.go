package main

import (
	"github.com/farelsukabelajar/task-5-pbi-btpns-KeyFirdausiAlfarel/database"
	"github.com/farelsukabelajar/task-5-pbi-btpns-KeyFirdausiAlfarel/router"
)

func main() {
	database.InitDB()
	database.MigrateDB()
	r := router.RouteInit()
	r.Run(":8080")
}
