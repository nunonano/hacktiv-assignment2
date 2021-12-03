package main

import (
	"github.com/nunonano/hacktiv-assignment2/app"
	"github.com/nunonano/hacktiv-assignment2/controller"
)

func main() {
	db := app.InitDB()
	DBConn := &controller.DBConn{DB: db}

	r := app.InitRouter(*DBConn)

	r.Run(":8080")
}

