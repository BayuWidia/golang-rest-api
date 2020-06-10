package main

import (
	"github.com/dudyali/rest-api/db"
	"github.com/dudyali/rest-api/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}
