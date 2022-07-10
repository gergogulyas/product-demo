package product

type Filter interface {
	ApplyOnItem(*Product) bool

	condition() func(p *Product) bool
}

type CategoryFilter struct {
	category string
}

func NewCategoryFilter(category string) *CategoryFilter {
	return &CategoryFilter{category: category}
}

func (f *CategoryFilter) condition() func(p *Product) bool {
	return func(p *Product) bool {
		return p.Category == f.category
	}
}

func (f *CategoryFilter) ApplyOnItem(product *Product) bool {
	valid := f.condition()

	return valid(product)
}

type PriceLessThanFilter struct {
	price int
}

func NewPriceLessThanFilter(price int) *PriceLessThanFilter {
	return &PriceLessThanFilter{price: price}
}

func (f *PriceLessThanFilter) condition() func(p *Product) bool {
	return func(p *Product) bool {
		return p.Price <= f.price
	}
}

func (f *PriceLessThanFilter) ApplyOnItem(product *Product) bool {
	valid := f.condition()

	return valid(product)
}
