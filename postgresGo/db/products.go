package db

import (
	"log"
	"time"

	"github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
)

// ProductItem consists of the attributes of our database
type ProductItem struct {
	tableName struct{} `sql:"product_items_collection"`
	ID        int      `sql:"id,pk"` //Primary key - pk
	Name      string   `sql:"name,unique"`
	Desc      string   `sql:"desc"`
	Image     string   `sql:"image"`
	Price     float64  `sql:"price,type:real"`
	Features  struct {
		Name string
		Desc string
		Imp  int
	} `sql:"features,type:jsonb"`
	CreatedAt time.Time `sql:"created_at"`
	UpdatedAt time.Time `sql:"updated_at"`
	IsActive  bool      `sql:"is_active"`
	// Note that Go types have to be converted into PostgreSQL types:
	// int8,uint8,int16  | smallint
	// uint16,int32,     | integer
	// uint32,int64,int  | bigint
	// uint, uint64      | bigint
	// float32           | real
	// float64           | double precision
	// bool              | boolean
	// string            | text
	// []byte            | bytea
	// struct,map,array  | jsonb
	// time.Time         | timestamptz
	// net.IP            | inet
	// net.IPNet         | cidr

}

// CreateProductItemsTable returns nil if the tables are successfully created
func CreateProductItemsTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&ProductItem{}, opts)
	if createErr != nil {
		log.Printf("Error while creating table productItems, Reason being: %v\n", createErr)
		return createErr
	}
	log.Printf("Table ProductItems created successfully.\n")
	return nil
}
