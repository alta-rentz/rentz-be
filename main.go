package main

import (
	"project3/config"
	"project3/routes"
)

func main() {
	config.InitDB()
	e := routes.New()

	e.Logger.Fatal(e.Start(":8080"))
}
