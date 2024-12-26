package main

import (
	"fmt"
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

	fmt.Println("Server at 9090")

	server := http.Server{
		Addr:    ":9090",
		Handler: handler.SecureHandler(r),
	}
	log.Fatal(server.ListenAndServe())

	//log.Fatal(http.ListenAndServe(":9090", handlers.CORS()(r)))
}
