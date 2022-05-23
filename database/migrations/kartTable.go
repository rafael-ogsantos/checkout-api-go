package migrations

import (
	"context"
	"log"

	"github.com/rafael-ogsantos/checkout-api-go/database"
	"github.com/rafael-ogsantos/checkout-api-go/models"
)

type KartTable struct {
}

func (k *KartTable) Migrate() bool {

	_, err2 := database.DB.NewCreateTable().
		Model((*models.Kart)(nil)).
		IfNotExists().
		ForeignKey(`("product_id") REFERENCES "products" ("id") ON DELETE CASCADE`).
		Exec(context.Background())

	if err2 != nil {
		log.Panic("Error to creation table", err2)
	}

	return true
}
