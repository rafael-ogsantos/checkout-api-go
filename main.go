package main

import (
	db "github.com/rafael-ogsantos/checkout-api-go/database"
	"github.com/rafael-ogsantos/checkout-api-go/database/migrations"
	"github.com/rafael-ogsantos/checkout-api-go/routes"
)

type Migration interface {
	Migrate() bool
}

func MigrateA(table Migration) {
	table.Migrate()
}

func main() {
	db.Connect()
	routes.HundleRequest()
	product := migrations.ProductTable{}
	kart := migrations.KartTable{}

	MigrateA(&product)
	MigrateA(&kart)
}
