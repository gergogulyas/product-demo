package product

import "strings"

type Discounter interface {
	Apply(*Product)
}

type DiscountRepository struct {
	Discounts []Discounter
}

func NewDiscountRepository() *DiscountRepository {
	repository := &DiscountRepository{}
	repository.init()

	return repository
}

func (r *DiscountRepository) init() {
	skuDiscount := NewSKUDiscount("000003", 15)
	categoryDiscount := NewCategoryDiscount("boots", 30)

	r.Discounts = append(r.Discounts, skuDiscount, categoryDiscount)
}

func NewCategoryDiscount(category string, percent int) CategoryDiscount {
	return CategoryDiscount{
		category: category,
		percent:  percent,
	}
}

type CategoryDiscount struct {
	category string
	percent  int
}

func (cd CategoryDiscount) Apply(product *Product) {
	if strings.EqualFold(product.Category, cd.category) && product.Discount() < cd.percent {
		product.SetDiscount(cd.percent)
	}
}

func NewSKUDiscount(sku string, percent int) SKUDiscount {
	return SKUDiscount{
		sku:     sku,
		percent: percent,
	}
}

type SKUDiscount struct {
	sku     string
	percent int
}

func (sd SKUDiscount) Apply(product *Product) {
	if strings.EqualFold(product.SKU, sd.sku) && product.Discount() < sd.percent {
		product.SetDiscount(sd.percent)
	}
}
