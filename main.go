package main

import (
	"log"

	"github.com/piyabch/pi-api/config"
	"github.com/piyabch/pi-api/db"
	"github.com/piyabch/pi-api/rest"
)

// The main function is the starting point of the application.
// The control flow are as follows.
//   - Load program configuration such as hostname, ports, and so on.
//   - Connect to the relational database.
//   - Start the rest server.
func main() {
	log.Print("Starting pi-api server")
	// load configurations
	config.LoadConfig("./config")
	// connect database
	db.Connect()
	// start rest server
	rest.Start()
}
