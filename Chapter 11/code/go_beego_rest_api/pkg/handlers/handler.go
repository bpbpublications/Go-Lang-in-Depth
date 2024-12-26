package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
)

type CustomerBody struct {
	Name    string `json"name"`
	Mobile  string `json"mobile"`
	Address string `json"address"`
}

type Customer struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Mobile  string `json:"mobile"`
	Address string `json:"address"`
}

func GetCustomers(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var customers []Customer
		_ = db.Table("customer").Select("id, name,mobile,address").Scan(&customers)
		json.NewEncoder(w).Encode(customers)
	}
}

func CreateCustomer(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var RequestBody CustomerBody
		json.NewDecoder(r.Body).Decode(&RequestBody)
		_ = db.Table("customer").Create(&RequestBody)
		fmt.Println("Created Customer")
		json.NewEncoder(w).Encode(RequestBody)
	}
}

func UpdateCustomer(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var PutBody Customer
		json.NewDecoder(r.Body).Decode(&PutBody)
		_ = db.Table("customer").Where("id=?", PutBody.Id).Update("name", PutBody.Name).Scan(&PutBody)
		fmt.Printf("Updated Customer with id %d\n", PutBody.Id)
		json.NewEncoder(w).Encode(PutBody)
	}
}

func DeleteCustomer(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var DeleteBody Customer
		json.NewDecoder(r.Body).Decode(&DeleteBody)
		_ = db.Table("customer").Delete(&DeleteBody)
		fmt.Printf("Deleted Customer with id %d\n", DeleteBody.Id)
		json.NewEncoder(w).Encode(DeleteBody)
	}
}
