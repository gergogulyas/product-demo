package product

func (r Repository) GetList() []Resource {
	list := []Resource{}

	for _, product := range r.Items.Products {
		for _, discountModel := range r.Discounts.Discounts {
			discountModel.Apply(&product)
		}

		list = append(list, Resource{
			SKU:      product.SKU,
			Name:     product.Name,
			Category: product.Category,
			Price:    product.getPriceResource(),
		})
	}

	return list
}
