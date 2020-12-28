package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book struct (Model)
type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"Name"`
	Title string  `json:"title"`
	Price float64 `json:"Price"`
}

var Item []Product

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint called: homePage()")

}
func getItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Item)
}
func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through books and find one with the id from the params
	for _, item := range Item {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Product{})
}
func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	Item = append(Item, product)
	json.NewEncoder(w).Encode(product)
}
func main() {
	// Init router
	r := mux.NewRouter()

	// Hardcoded data - @todo: add database
	Item = append(Item, Product{
		ID:    "1",
		Name:  "Rice",
		Title: "Product One",
		Price: 65})
	Item = append(Item, Product{
		ID:    "2",
		Name:  "Apple",
		Title: "Product Two",
		Price: 126})

	// Route handles & endpoints
	r.HandleFunc("/", HomePage).Methods("GET")
	r.HandleFunc("/Item", getItem).Methods("GET")
	r.HandleFunc("/Item/{id}", getProduct).Methods("GET")
	r.HandleFunc("/Item", createItem).Methods("POST")
	// Start server
	log.Fatal(http.ListenAndServe(":5050", r))
}
