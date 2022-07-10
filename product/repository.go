package product

import (
	"encoding/json"
	"io/ioutil"
)

type RepositoryContract interface {
	GetList(filters []Filter) []Resource
}

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
	repository := &Repository{Discounts: discounts}
	repository.init()

	return repository
}
