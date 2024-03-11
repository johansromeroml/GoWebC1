package main

import (
	"class1api/internal/handler"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	//fmt.Println(products.Products[4])
	ph := handler.NewProductHandler()

	mux := chi.NewRouter()

	mux.Get("/ping", ph.Pong)

	mux.Route("/products", func(router chi.Router) {
		// get all products
		router.Get("/", ph.ListAllProducts)
		// get product by id
		//rt.Get("/{id}", hd.GetByID())
		// search products
		//rt.Get("/search", hd.Search())
		// create product
		//rt.Post("/", hd.Create())
		// update or create product
		//rt.Put("/{id}", hd.UpdateOrCreate())
		// update product
		//rt.Patch("/{id}", hd.Update())
		// delete product
		//rt.Delete("/{id}", hd.Delete())
	})

	fmt.Println("Starting server ...")
	//fmt.Println(mux)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
