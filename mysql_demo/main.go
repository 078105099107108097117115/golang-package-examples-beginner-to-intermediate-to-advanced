package main

import (
	// "fmt"
	// "github.com/go-sql-driver/mysql"
	"database/sql"
)

func main() {
	db, err := sql.Open()
	// Open opens a database specified by its database driver name
	// and a driver-specific data source name,
	// usually consisting of at least a database
	// name and connection information.
	// Most users will open a database via a driver-specific connection helper
	// function that returns a *DB. No database drivers are included in the Go
	// standard library
	//DSN(Data source name):->[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	// Except for the databasename, all values are optional. So the minimal DSN is:
	// /dbname
}
