package main

import (
	"log"

	"github.com/piyabch/pi-api/config"
	"github.com/piyabch/pi-api/db"
	"github.com/piyabch/pi-api/rest"
)

func main() {
	log.Print("Starting pi-api server")
	config.LoadConfig("./config")
	db.Connect()
	rest.Init()
	rest.Start()
}
