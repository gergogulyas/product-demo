package product

import (
	"encoding/json"
	"io/ioutil"
)

type Repository struct {
	Discounts *DiscountRepository
	Items     List
}

func (r *Repository) init() {
	data := List{}
	file, _ := ioutil.ReadFile("product.json")
	_ = json.Unmarshal(file, &data)

	r.Items = data
}

func NewProductRepository(discounts *DiscountRepository) *Repository {
	repo := &Repository{Discounts: discounts}
	repo.init()

	return repo
}
