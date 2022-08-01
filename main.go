package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	db2 "github.com/raphaelmb/fullcycle-arch-hexagonal/adapters/db"
	"github.com/raphaelmb/fullcycle-arch-hexagonal/application"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	product, _ := productService.Create("Product", 30)

	productService.Enable(product)
}
