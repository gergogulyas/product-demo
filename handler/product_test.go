package handler_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/gergogulyas/product-demo/handler"
	"github.com/gergogulyas/product-demo/product"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (m *ProductRepositoryMock) GetList(filters []product.Filter) []product.Resource {
	args := m.Called(filters)
	mockResponse := args.Get(0).([]product.Resource)

	return mockResponse
}

type ProductHandlerTestSuite struct {
	ProductRepositoryMock *ProductRepositoryMock

	suite.Suite
}

func TestProductHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(ProductHandlerTestSuite))
}

func (s *ProductHandlerTestSuite) SetupTest() {
	s.ProductRepositoryMock = new(ProductRepositoryMock)
}

func (s *ProductHandlerTestSuite) getProductEndpointResponse(categoryParam *string, priceLessThanParam *int) (*http.Response, []byte) {
	h := handler.Product{Repository: s.ProductRepositoryMock}

	uri := "/products?"
	if categoryParam != nil {
		uri += fmt.Sprintf("category=%s&", *categoryParam)
	}
	if priceLessThanParam != nil {
		uri += fmt.Sprintf("priceLessThan=%d", *priceLessThanParam)
	}

	req := httptest.NewRequest("GET", uri, nil)
	w := httptest.NewRecorder()
	h.List(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	return resp, body
}

func (s *ProductHandlerTestSuite) TestRepositoryGetListCalledAnd200ResponseCodeReturned() {
	filters := []product.Filter{}
	s.ProductRepositoryMock.On("GetList", filters).Once().Return([]product.Resource{})

	resp, body := s.getProductEndpointResponse(nil, nil)

	s.Assert().Equal(http.StatusOK, resp.StatusCode)
	s.Assert().JSONEq("[]", string(body))
}

func (s *ProductHandlerTestSuite) TestCategoryFilterPassed() {
	filters := []product.Filter{product.NewCategoryFilter("testing")}

	s.ProductRepositoryMock.On("GetList", filters).Once().Return([]product.Resource{})

	categoryParam := "testing"
	resp, body := s.getProductEndpointResponse(&categoryParam, nil)

	s.Assert().Equal(http.StatusOK, resp.StatusCode)
	s.Assert().JSONEq("[]", string(body))

	s.ProductRepositoryMock.AssertCalled(s.T(), "GetList", filters)
	s.ProductRepositoryMock.AssertNumberOfCalls(s.T(), "GetList", 1)
}

func (s *ProductHandlerTestSuite) TestPriceLessThanFilterPassed() {
	filters := []product.Filter{product.NewPriceLessThanFilter(10000)}

	s.ProductRepositoryMock.On("GetList", filters).Once().Return([]product.Resource{})

	priceLessThanParam := 10000
	resp, body := s.getProductEndpointResponse(nil, &priceLessThanParam)

	s.Assert().Equal(http.StatusOK, resp.StatusCode)
	s.Assert().JSONEq("[]", string(body))

	s.ProductRepositoryMock.AssertCalled(s.T(), "GetList", filters)
	s.ProductRepositoryMock.AssertNumberOfCalls(s.T(), "GetList", 1)
}

func (s *ProductHandlerTestSuite) TestBothFilterPassed() {
	filters := []product.Filter{
		product.NewCategoryFilter("testing"),
		product.NewPriceLessThanFilter(10000),
	}

	s.ProductRepositoryMock.On("GetList", filters).Once().Return([]product.Resource{})

	categoryParam := "testing"
	priceLessThanParam := 10000
	resp, body := s.getProductEndpointResponse(&categoryParam, &priceLessThanParam)

	s.Assert().Equal(http.StatusOK, resp.StatusCode)
	s.Assert().JSONEq("[]", string(body))

	s.ProductRepositoryMock.AssertCalled(s.T(), "GetList", filters)
	s.ProductRepositoryMock.AssertNumberOfCalls(s.T(), "GetList", 1)
}

func (s *ProductHandlerTestSuite) TestReturnExpectedItemsWithoutDiscount() {
	filters := []product.Filter{}
	expectedResponse := []product.Resource{
		{
			SKU:      "000001",
			Name:     "Test name",
			Category: "Test category",
			Price: product.PriceResource{
				Original:           10000,
				Final:              10000,
				DiscountPercentage: nil,
				Currency:           "EUR",
			},
		},
	}

	s.ProductRepositoryMock.On("GetList", filters).Once().Return(expectedResponse)

	resp, body := s.getProductEndpointResponse(nil, nil)

	s.Assert().Equal(http.StatusOK, resp.StatusCode)
	responseJson, _ := json.Marshal(expectedResponse)
	s.Assert().JSONEq(string(responseJson), string(body))

	s.ProductRepositoryMock.AssertCalled(s.T(), "GetList", filters)
	s.ProductRepositoryMock.AssertNumberOfCalls(s.T(), "GetList", 1)
}

func (s *ProductHandlerTestSuite) TestReturnExpectedItemsWithDiscount() {
	filters := []product.Filter{}
	discountText := "10%"
	expectedResponse := []product.Resource{
		{
			SKU:      "000001",
			Name:     "Test name",
			Category: "Test category",
			Price: product.PriceResource{
				Original:           10000,
				Final:              9000,
				DiscountPercentage: &discountText,
				Currency:           "EUR",
			},
		},
	}

	s.ProductRepositoryMock.On("GetList", filters).Once().Return(expectedResponse)

	resp, body := s.getProductEndpointResponse(nil, nil)

	s.Assert().Equal(http.StatusOK, resp.StatusCode)
	responseJson, _ := json.Marshal(expectedResponse)
	s.Assert().JSONEq(string(responseJson), string(body))

	s.ProductRepositoryMock.AssertCalled(s.T(), "GetList", filters)
	s.ProductRepositoryMock.AssertNumberOfCalls(s.T(), "GetList", 1)
}
