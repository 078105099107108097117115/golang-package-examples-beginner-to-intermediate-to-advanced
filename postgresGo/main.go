package main

import (
	"fmt"

	db "github.com/nicklausnn/postgresGo/db"
)

func main() {
	fmt.Println("PostgreSQL and Golang")
	db.Connect()

}
