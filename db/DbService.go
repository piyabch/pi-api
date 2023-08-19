package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/piyabch/pi-api/config"
)

var db *sql.DB

// Connect to MySQL database using configuration Struct
func Connect() {
	log.Print("Connecting to pi_data DB")
	// make MySQL configuration Struct
	cfg := mysql.Config{
		User:   config.App.MySqlUsername,
		Passwd: config.App.MySqlPassword,
		Net:    "tcp",
		Addr:   config.App.MySqlAddr,
		DBName: config.App.MySqlDbName,
	}
	// open a database connection and keep the driver reference
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	// check whether the connection is success
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	log.Print("pi_data DB connected")
}
