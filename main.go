package main

import (
	db "github.com/rafael-ogsantos/checkout-api-go/database"
	"github.com/rafael-ogsantos/checkout-api-go/database/migrations"
)

type Migration interface {
	Migrate() bool
}

func MigrateA(table Migration) {
	table.Migrate()
}

func main() {
	db.Connect()
	// product := migrations.ProductTable{}
	kart := migrations.KartTable{}

	// MigrateA(&product)
	MigrateA(&kart)

	// product.Migrate()
	kart.Migrate()
}
