package main

import (
	"project-ecommerce/config"
	"project-ecommerce/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
