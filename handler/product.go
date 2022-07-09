package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gergogulyas/product-demo/product"
)

type Product struct {
	Repository *product.Repository
}

func (h Product) List(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var filters []product.Filter
	if req.URL.Query().Has("category") {
		c := req.URL.Query().Get("category")
		filters = append(filters, product.NewCategoryFilter(c))
	}

	l := h.Repository.GetList(filters)
	_ = json.NewEncoder(w).Encode(l)
}
