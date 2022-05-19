package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/rafael-ogsantos/checkout-api-go/database"
	"github.com/rafael-ogsantos/checkout-api-go/models"
)

func AllProducts(w http.ResponseWriter, r *http.Request) {
	var p []models.Product
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}
