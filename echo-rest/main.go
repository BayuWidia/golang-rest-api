package main

import (
	"github.com/bayuwidia/echo-rest/db"
	"github.com/bayuwidia/echo-rest/routes"
)

func main() {

	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}
