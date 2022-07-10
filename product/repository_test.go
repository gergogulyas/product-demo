package product_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/gergogulyas/product-demo/product"
)

type RepositoryTestSuite struct {
	suite.Suite
}

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (s *RepositoryTestSuite) TestNewRepositoryReturnExpectedType() {
	discountRepository := product.DiscountRepository{}
	repo := product.NewProductRepository(&discountRepository)

	s.Assert().IsType(&product.Repository{}, repo)
	s.Assert().Implements((*product.RepositoryContract)(nil), repo)
}
