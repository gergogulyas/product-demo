package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gergogulyas/product-demo/product"
)

type Product struct {
	Repository product.RepositoryContract
}

func (h Product) List(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	filters := []product.Filter{}
	if req.URL.Query().Has("category") {
		c := req.URL.Query().Get("category")
		filters = append(filters, product.NewCategoryFilter(c))
	}

	if req.URL.Query().Has("priceLessThan") {
		parsePrice, err := strconv.ParseInt(req.URL.Query().Get("priceLessThan"), 0, 32)
		if err == nil {
			filters = append(filters, product.NewPriceLessThanFilter(int(parsePrice)))
		}
	}

	l := h.Repository.GetList(filters)
	_ = json.NewEncoder(w).Encode(l)
}
