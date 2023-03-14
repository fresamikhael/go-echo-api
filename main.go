package main

import (
	"api-echo-go/db"
	"api-echo-go/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1323"))
}
