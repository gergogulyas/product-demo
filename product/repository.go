package product

import (
	"encoding/json"
	"io/ioutil"
)

type Repository struct {
	Discounts *DiscountRepository
	Items     []Product
}

func (r *Repository) init() {
	data := List{}
	file, _ := ioutil.ReadFile("product.json")
	_ = json.Unmarshal(file, &data)

	r.Items = data.Products
}

func NewProductRepository(discounts *DiscountRepository) *Repository {
	repo := &Repository{Discounts: discounts}
	repo.init()

	return repo
}
