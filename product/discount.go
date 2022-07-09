package product

import "strings"

type DiscountRepository struct {
	Discounts []Discounter
}

func (r *DiscountRepository) init() {
	skuDiscount := SKUDiscount{SKU: "000003", Percent: 15}
	categoryDiscount := CategoryDiscount{Category: "boots", Percent: 30}

	r.Discounts = append(r.Discounts, skuDiscount, categoryDiscount)
}

func NewDiscountRepository() *DiscountRepository {
	discountR := &DiscountRepository{}
	discountR.init()

	return discountR
}

type Discounter interface {
	Apply(*Product)
}

type CategoryDiscount struct {
	Category string
	Percent  int
}

func (cd CategoryDiscount) Apply(product *Product) {
	if strings.EqualFold(product.Category, cd.Category) && product.Discount() < cd.Percent {
		product.SetDiscount(cd.Percent)
	}
}

type SKUDiscount struct {
	SKU     string
	Percent int
}

func (sd SKUDiscount) Apply(product *Product) {
	if strings.EqualFold(product.SKU, sd.SKU) && product.Discount() < sd.Percent {
		product.SetDiscount(sd.Percent)
	}
}
