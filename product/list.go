package product

const maxItems = 5

func (r Repository) GetList(filters []Filter) []Resource {
	list := []Resource{}

	listCounter := 0
	for _, product := range r.Items {
		isValid := true

		for _, filter := range filters {
			isValid = filter.ApplyOnItem(&product)
			if !isValid {
				break // stop iteration on first false
			}
		}

		if !isValid {
			continue // goto next item
		}

		for _, discountModel := range r.Discounts.Discounts {
			discountModel.Apply(&product)
		}

		list = append(list, Resource{
			SKU:      product.SKU,
			Name:     product.Name,
			Category: product.Category,
			Price:    product.getPriceResource(),
		})

		listCounter++
		if listCounter >= maxItems {
			break
		}
	}

	return list
}
