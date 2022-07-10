package product_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/gergogulyas/product-demo/product"
)

type FilterTestSuite struct {
	suite.Suite
}

func TestFilterTestSuite(t *testing.T) {
	suite.Run(t, new(FilterTestSuite))
}

func (s *FilterTestSuite) TestCategoryFilterAppliedOnProduct() {
	filter := product.NewCategoryFilter("test")
	p := product.Product{
		SKU:      "000001",
		Name:     "Test Name",
		Category: "test",
		Price:    10000,
	}
	filter.ApplyOnItem(&p)

	s.Assert().True(filter.ApplyOnItem(&p))
}

func (s *FilterTestSuite) TestCategoryFilterNotAppliedOnProduct() {
	filter := product.NewCategoryFilter("test")
	p := product.Product{
		SKU:      "000001",
		Name:     "Test Name",
		Category: "Non test",
		Price:    10000,
	}

	s.Assert().False(filter.ApplyOnItem(&p))
}

func (s *FilterTestSuite) TestPriceLessThanFilterAppliedOnProduct() {
	filter := product.NewPriceLessThanFilter(10000)
	p := product.Product{
		SKU:      "000001",
		Name:     "Test Name",
		Category: "test",
		Price:    10000,
	}
	filter.ApplyOnItem(&p)

	s.Assert().True(filter.ApplyOnItem(&p))
}

func (s *FilterTestSuite) TestPriceLessThanFilterNotAppliedOnProduct() {
	filter := product.NewPriceLessThanFilter(9999)
	p := product.Product{
		SKU:      "000001",
		Name:     "Test Name",
		Category: "Non test",
		Price:    10000,
	}

	s.Assert().False(filter.ApplyOnItem(&p))
}
