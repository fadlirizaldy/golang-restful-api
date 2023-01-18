package main

import (
	"project_alterra/config"
	"project_alterra/router"
)



func main() {

	// Initansiasi DB
	config.InitDB()

	//Instansiasi server
	e := router.New()
	

	e.Logger.Fatal(e.Start(":1323"))
}