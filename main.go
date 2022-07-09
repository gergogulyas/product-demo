package main

import (
	"net/http"

	"github.com/gergogulyas/product-demo/handler"
	"github.com/gergogulyas/product-demo/product"
)

func main() {
	discountRepository := product.NewDiscountRepository()
	productRepository := product.NewProductRepository(discountRepository)

	productHandler := handler.Product{Repository: productRepository}

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			productHandler.List(w, r)
		} else {
			http.Error(w, "Invalid request method.", 405)
		}
	})

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
}
