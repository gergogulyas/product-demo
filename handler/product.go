package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gergogulyas/product-demo/product"
)

type Product struct {
	Repository *product.Repository
}

func (h Product) List(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	l := h.Repository.GetList()
	_ = json.NewEncoder(w).Encode(l)
}
