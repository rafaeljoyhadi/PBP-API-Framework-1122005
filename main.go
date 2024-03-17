package main

import (
	"echo-api/routes"
)

func main() {
	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}
