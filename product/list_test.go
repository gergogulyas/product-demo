package product_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/gergogulyas/product-demo/product"
)

type ListTestSuite struct {
	suite.Suite
}

func TestListTestSuite(t *testing.T) {
	suite.Run(t, new(ListTestSuite))
}

func (s *ModelTestSuite) TestGetListReturnEmptySliceWhenIteratingOnEmptySlice() {
	r := product.NewProductRepository(&product.DiscountRepository{})
	r.Items = []product.Product{}
	result := r.GetList([]product.Filter{})

	s.Assert().Empty(result)
}

func (s *ModelTestSuite) TestGetListReturnMaximumFiveItems() {
	r := product.NewProductRepository(&product.DiscountRepository{})
	r.Items = []product.Product{
		{},
		{},
		{},
		{},
		{},
		{},
		{},
	}

	result := r.GetList([]product.Filter{})

	s.Assert().NotEmpty(result)
	s.Assert().Equal(5, len(result))
}

func (s *ModelTestSuite) TestGetListReturnExpectedItemsWhenFiltersPassed() {
	r := product.NewProductRepository(&product.DiscountRepository{})
	r.Items = []product.Product{
		{Category: "test", Price: 15000},
		{Category: "test", Price: 30000},
		{Category: "test", Price: 20000},
		{Category: "non test", Price: 20000},
	}

	result := r.GetList([]product.Filter{
		product.NewPriceLessThanFilter(20000),
		product.NewCategoryFilter("test"),
	})

	s.Assert().NotEmpty(result)
	s.Assert().Equal(2, len(result))
}

func (s *ModelTestSuite) TestGetListReturnItemWithDiscountApplied() {
	cd := product.NewCategoryDiscount("test", 10)
	sd := product.NewSKUDiscount("000001", 20)
	dr := product.DiscountRepository{}
	dr.Discounts = []product.Discounter{cd, sd}

	r := product.NewProductRepository(&dr)
	r.Items = []product.Product{
		{SKU: "000001", Category: "test", Price: 10000},
		{SKU: "000002", Category: "non test", Price: 20000},
		{SKU: "000003", Category: "test", Price: 30000},
		{SKU: "000001", Category: "non test", Price: 40000},
	}

	result := r.GetList([]product.Filter{})

	discountTenPercentText := "10%"
	discountTwentyPercentText := "20%"

	expectedResult := []product.Resource{
		{
			SKU:      "000001",
			Category: "test",
			Name:     "",
			Price: product.PriceResource{
				Original:           10000,
				Final:              8000,
				DiscountPercentage: &discountTwentyPercentText,
				Currency:           "EUR",
			},
		},
		{
			SKU:      "000002",
			Category: "non test",
			Name:     "",
			Price: product.PriceResource{
				Original:           20000,
				Final:              20000,
				DiscountPercentage: nil,
				Currency:           "EUR",
			},
		},
		{
			SKU:      "000003",
			Category: "test",
			Name:     "",
			Price: product.PriceResource{
				Original:           30000,
				Final:              27000,
				DiscountPercentage: &discountTenPercentText,
				Currency:           "EUR",
			},
		},
		{
			SKU:      "000001",
			Category: "non test",
			Name:     "",
			Price: product.PriceResource{
				Original:           40000,
				Final:              32000,
				DiscountPercentage: &discountTwentyPercentText,
				Currency:           "EUR",
			},
		},
	}

	s.Assert().Equal(expectedResult, result)
}
