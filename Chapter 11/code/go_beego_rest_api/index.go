package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"go_beego_rest_api/pkg/db"
	handler "go_beego_rest_api/pkg/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	DB := db.InitializeDB()

	r.HandleFunc("/customers", handler.GetCustomers(DB)).Methods("GET")
	r.HandleFunc("/create", handler.CreateCustomer(DB)).Methods("POST")
	r.HandleFunc("/update", handler.UpdateCustomer(DB)).Methods("PUT")
	r.HandleFunc("/delete", handler.DeleteCustomer(DB)).Methods("DELETE")

	fmt.Println("Server at 8080")

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS()(r)))
}
