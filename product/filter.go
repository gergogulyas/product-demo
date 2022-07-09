package product

type Filter interface {
	ApplyOnList([]Product) []Product
	ApplyOnItem(Product) bool

	condition() func(p Product) bool
}

type CategoryFilter struct {
	category string
}

func NewCategoryFilter(category string) *CategoryFilter {
	return &CategoryFilter{category: category}
}

func (f *CategoryFilter) condition() func(p Product) bool {
	return func(p Product) bool {
		return p.Category == f.category
	}
}

func (f *CategoryFilter) ApplyOnItem(product Product) bool {
	valid := f.condition()

	return valid(product)
}

func (f *CategoryFilter) ApplyOnList(list []Product) []Product {
	filtered := []Product{}
	valid := f.condition()

	for _, product := range list {
		if valid(product) {
			filtered = append(filtered, product)
		}
	}

	return filtered
}

type PriceLessThanFilter struct {
	price int
}

func NewPriceLessThanFilter(price int) *PriceLessThanFilter {
	return &PriceLessThanFilter{price: price}
}

func (f *PriceLessThanFilter) condition() func(p Product) bool {
	return func(p Product) bool {
		return p.Price <= f.price
	}
}

func (f *PriceLessThanFilter) ApplyOnItem(product Product) bool {
	valid := f.condition()

	return valid(product)
}

func (f *PriceLessThanFilter) ApplyOnList(list []Product) []Product {
	filtered := []Product{}
	valid := f.condition()

	for _, product := range list {
		if valid(product) {
			filtered = append(filtered, product)
		}
	}

	return filtered
}
