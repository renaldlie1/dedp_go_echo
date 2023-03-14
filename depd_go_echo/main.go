package main

import (
	"github.com/renaldlie1/depd_go_echo/db"
	"github.com/renaldlie1/depd_go_echo/routes"
)

func main() {
	db.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":9999"))
}
