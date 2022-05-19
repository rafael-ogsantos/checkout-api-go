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

	_, err2 := database.DB.NewCreateTable().
		Model((*models.Product)(nil)).
		Exec(context.Background())

	if err2 != nil {
		log.Panic("Error to creation table", err2)
		return false
	}

	return true
}
