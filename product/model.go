package product

import "fmt"

type List struct {
	Products []Product `json:"products"`
}

type Product struct {
	SKU             string `json:"sku"`
	Name            string `json:"name"`
	Category        string `json:"category"`
	Price           int    `json:"price"`
	discountPercent int
}

func (p Product) HasDiscount() bool {
	return p.discountPercent != 0
}

func (p Product) Discount() int {
	return p.discountPercent
}

func (p *Product) SetDiscount(percent int) {
	p.discountPercent = percent
}

func (p Product) getPriceResource() PriceResource {
	if p.HasDiscount() {
		percentText := fmt.Sprintf("%d%%", p.Discount())
		d := (100.0 - float64(p.Discount())) / 100.0
		return PriceResource{
			Original:           p.Price,
			Final:              int(float64(p.Price) * d),
			DiscountPercentage: &percentText,
			Currency:           "EUR",
		}
	}

	return PriceResource{
		Original:           p.Price,
		Final:              p.Price,
		DiscountPercentage: nil,
		Currency:           "EUR",
	}

}
