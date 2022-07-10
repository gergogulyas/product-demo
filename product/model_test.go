package product_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/gergogulyas/product-demo/product"
)

type ModelTestSuite struct {
	suite.Suite
}

func TestModelTestSuite(t *testing.T) {
	suite.Run(t, new(ModelTestSuite))
}

func (s *ModelTestSuite) TestHasDiscountReturnsFalseWhenNoDiscountApplied() {
	p := product.Product{}

	s.Assert().False(p.HasDiscount())
}

func (s *ModelTestSuite) TestHasDiscountReturnsTrueAfterSetDiscount() {
	p := product.Product{}
	p.SetDiscount(10)

	s.Assert().True(p.HasDiscount())
}

func (s *ModelTestSuite) TestDiscountReturnsExpectedValueAfterSetDiscount() {
	p := product.Product{}
	p.SetDiscount(10)

	s.Assert().Equal(p.Discount(), 10)
}
