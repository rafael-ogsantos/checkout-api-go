package migrations

import (
	"context"
	"log"

	"github.com/rafael-ogsantos/checkout-api-go/database"
	"github.com/rafael-ogsantos/checkout-api-go/models"
)

type ProductTable struct {
}

func (p *ProductTable) Migrate() bool {
	ctx := context.Background()
	_, err := database.DB.NewCreateTable().
		IfNotExists().
		Model((*models.Product)(nil)).
		Exec(ctx)

	if err != nil {
		log.Panic("Error to creation table ", err)
		return false
	}

	product1 := models.Product{Name: "Jogger preta", Category: "Calça", Quantity: 4}
	product2 := models.Product{Name: "Camisa Nike M", Category: "Camiseta", Quantity: 10}
	product3 := models.Product{Name: "Tênis New Balance", Category: "Tênis", Quantity: 5}
	products := []models.Product{product1, product2, product3}

	_, err = database.DB.NewInsert().Model(&products).Exec(ctx)
	if err != nil {
		panic(err)
	}

	return true
}
