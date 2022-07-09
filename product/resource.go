package product

type PriceResource struct {
	Original           int     `json:"original"`
	Final              int     `json:"final"`
	DiscountPercentage *string `json:"discount_percentage"` // FIXME omitempty
	Currency           string  `json:"currency"`
}

type Resource struct {
	SKU      string        `json:"sku"`
	Name     string        `json:"name"`
	Category string        `json:"category"`
	Price    PriceResource `json:"price"`
}
