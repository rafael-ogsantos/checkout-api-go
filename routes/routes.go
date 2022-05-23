package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rafael-ogsantos/checkout-api-go/controllers"
)

func HundleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Kart)
	log.Fatal(http.ListenAndServe(":8000", r))
}
