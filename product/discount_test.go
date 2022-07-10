package product_test

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/gergogulyas/product-demo/product"
)

type DiscountTestSuite struct {
	suite.Suite
}

func TestDiscountTestSuite(t *testing.T) {
	suite.Run(t, new(DiscountTestSuite))
}

func (s *DiscountTestSuite) TestNewDiscountRepositoryReturnExpectedTypeAndInitialized() {
	repo := product.NewDiscountRepository()

	s.Assert().IsType(&product.DiscountRepository{}, repo)
	s.Assert().NotEmpty(repo.Discounts)
}

func (s *DiscountTestSuite) TestCategoryFilterAppliedOnProduct() {
	randomPercent := rand.Int()
	discount := product.NewCategoryDiscount("test", randomPercent)
	p := product.Product{
		Category: "test",
	}
	discount.Apply(&p)

	s.Assert().True(p.HasDiscount())
	s.Assert().Equal(p.Discount(), randomPercent)
}

func (s *DiscountTestSuite) TestCategoryFilterNotAppliedOnProduct() {
	discount := product.NewCategoryDiscount("test", rand.Int())
	p := product.Product{
		Category: "Non test",
	}
	discount.Apply(&p)

	s.Assert().False(p.HasDiscount())
}

func (s *DiscountTestSuite) TestSKUFilterAppliedOnProduct() {
	randomPercent := rand.Int()
	discount := product.NewSKUDiscount("000001", randomPercent)
	p := product.Product{
		SKU: "000001",
	}
	discount.Apply(&p)

	s.Assert().True(p.HasDiscount())
	s.Assert().Equal(p.Discount(), randomPercent)
}

func (s *DiscountTestSuite) TestSKUFilterNotAppliedOnProduct() {
	discount := product.NewSKUDiscount("000002", rand.Int())
	p := product.Product{
		SKU: "000001",
	}
	discount.Apply(&p)

	s.Assert().False(p.HasDiscount())
}

func (s *DiscountTestSuite) TestBiggestDiscountAppliedOnProduct() {
	skuDiscount := product.NewSKUDiscount("000001", 15)
	categoryDiscount := product.NewCategoryDiscount("test", 30)

	p := product.Product{
		SKU:      "000001",
		Category: "test",
	}

	skuDiscount.Apply(&p)
	categoryDiscount.Apply(&p)

	s.Assert().True(p.HasDiscount())
	s.Assert().Equal(p.Discount(), 30)
}
