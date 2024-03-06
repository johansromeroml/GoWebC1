package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Product struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Quantity       int     `json:"quantity"`
	CodeValue      string  `json:"code_value"`
	Published      bool    `json:"is_published"`
	ExpirationDate string  `json:"expiration"`
	Price          float64 `json:"price"`
}

type ProductList struct {
	Products []Product
}

var products ProductList

func main() {
	productsFile, err := os.Open("products.json")

	if err != nil {
		println(err.Error())
	}
	defer func(f *os.File) {
		f.Close()
		fmt.Println("File closed")
	}(productsFile)

	productsJSON, err := io.ReadAll(productsFile)

	json.Unmarshal(productsJSON, &products.Products)

	fmt.Println(products.Products[5])

	if err != nil {
		println(err.Error())
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/ping", ping)

	mux.HandleFunc("/products", productList)

	mux.HandleFunc("/products/{id}", productByID)

	mux.HandleFunc("/products/search?value=", searchProductsPriceGreatherThanValue)

	fmt.Println("Starting server ...")
	//fmt.Println(mux)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func ping(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "pong")
}

func productList(w http.ResponseWriter, req *http.Request) {
	responseJSON, err := json.Marshal(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
	}
	fmt.Fprint(w, string(responseJSON))
}

func productByID(w http.ResponseWriter, req *http.Request) {
	idParam := req.PathValue("id")
	fmt.Fprint(w, idParam)
}

func searchProductsPriceGreatherThanValue(w http.ResponseWriter, req *http.Request) {
	valueParam := req.URL.Query().Get("value")
	value, err := strconv.ParseFloat(valueParam, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	result := productsPriceGreaterThanValue(value)
	responseJSON, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
	}
	fmt.Fprint(w, string(responseJSON))
}

func productsPriceGreaterThanValue(value float64) ProductList {
	var returnList ProductList
	for _, prod := range products.Products {
		if prod.Price > value {
			returnList.Products = append(returnList.Products, prod)
		}
	}
	return returnList
}

//mux.HandleFunc("/greetings", greetings)

/*

type Person struct {
	FirstName string
	LastName  string
}

func greetings(w http.ResponseWriter, req *http.Request) {
	var per Person
	payload, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if err := json.Unmarshal(payload, &per); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Fprintf(w, "Hello %s %s", per.FirstName, per.LastName)
}

*/
