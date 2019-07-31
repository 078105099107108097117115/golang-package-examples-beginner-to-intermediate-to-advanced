package db

import (
	"log"
	"os"

	pg "github.com/go-pg/pg"
)

//Connect initializes the connection to the Postgres database from our main function
func Connect() {
	opts := &pg.Options{
		User:     "nick",
		Password: "ganymede",
		Addr:     "localhost:5432",
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect to the database.\n")
		os.Exit(1)
	}
	log.Printf("Connection to database successful!\n")
	CreateProductItemsTable(db)
	closeErr := db.Close()
	if closeErr != nil {
		log.Printf("Error while closing the database connection. %v\n", closeErr)
		os.Exit(100)
	}
	log.Printf("Connection closed successfully")
}
