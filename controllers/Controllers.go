package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/rafael-ogsantos/checkout-api-go/database"
	"github.com/rafael-ogsantos/checkout-api-go/models"
)

// func AllProducts(w http.ResponseWriter, r *http.Request) {
// 	var p []models.Product
// 	database.DB.(&p)
// 	json.NewEncoder(w).Encode(p)
// }

func Kart(w http.ResponseWriter, r *http.Request) {
	var karts []models.Kart
	// kart := new(models.Kart)
	err := database.DB.NewSelect().
		Model(&karts).
		// ColumnExpr("*").
		Relation("Product").
		// Join("JOIN products as p ON p.id = kart.product_id").
		Scan(context.Background())

	if err != nil {
		println("Erro ", err)
	}
	json.NewEncoder(w).Encode(karts)
}

// ColumnExpr("*").
// 		Join("JOIN products as p ON p.id = kart.product_id").
