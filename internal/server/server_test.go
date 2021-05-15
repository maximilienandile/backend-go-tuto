package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/maximilienandile/backend-go-tuto/internal/category"

	"github.com/maximilienandile/backend-go-tuto/internal/uniqueid"

	"github.com/maximilienandile/backend-go-tuto/internal/storage"

	"github.com/golang/mock/gomock"

	"github.com/maximilienandile/backend-go-tuto/internal/product"

	"github.com/stretchr/testify/assert"
)

func TestServer_CreateProduct(t *testing.T) {
	// GIVEN
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockedStorage := storage.NewMockStorage(ctrl)
	mockedUniqueIDGenerator := uniqueid.NewMockGenerator(ctrl)
	testServer, err := New(Config{
		Port:              8080,
		Storage:           mockedStorage,
		UniqueIDGenerator: mockedUniqueIDGenerator,
	})
	assert.NoError(t, err, "building a server should not return an error")
	recorder := httptest.NewRecorder()
	inputProduct := product.Product{
		Name: "test product",
	}
	inputProductJSON, err := json.Marshal(inputProduct)
	assert.NoError(t, err, "should be able to marshall product")
	req, err := http.NewRequest("POST", "/admin/products", bytes.NewReader(inputProductJSON))
	assert.NoError(t, err, "no error should happend when building the request")
	req.Header.Add("Authorization", "ABC")
	// mocks expectations
	mockedStorage.EXPECT().CreateProduct(gomock.Any()).Return(nil)
	mockedUniqueIDGenerator.EXPECT().Generate().Return("foo")
	// WHEN
	testServer.Engine.ServeHTTP(recorder, req)

	// THEN
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "{\"id\":\"foo\",\"name\":\"test product\",\"image\":\"\",\"shortDescription\":\"\",\"description\":\"\",\"priceVatExcluded\":{\"money\":null,\"display\":\"\"},\"vat\":{\"money\":null,\"display\":\"\"},\"totalPrice\":{\"money\":null,\"display\":\"\"},\"stock\":0,\"reserved\":0,\"version\":0}", recorder.Body.String())
}

func TestServer_CreateCategory(t *testing.T) {
	// GIVEN
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockedStorage := storage.NewMockStorage(ctrl)
	mockedUniqueIDGenerator := uniqueid.NewMockGenerator(ctrl)
	testServer, err := New(Config{
		Port:              8080,
		Storage:           mockedStorage,
		UniqueIDGenerator: mockedUniqueIDGenerator,
	})
	assert.NoError(t, err, "building a server should not return an error")
	recorder := httptest.NewRecorder()

	categoryInput := category.Category{
		Name:        "Test name",
		Description: "Test Description",
	}
	categoryInputJSON, err := json.Marshal(categoryInput)
	assert.NoError(t, err, "impossible to marshall category")
	req, err := http.NewRequest("POST", "/admin/categories", bytes.NewReader(categoryInputJSON))
	assert.NoError(t, err, "no error should happend when building the request")
	req.Header.Add("Authorization", "ABC")
	// mocks expectations
	mockedUniqueIDGenerator.EXPECT().Generate().Return("foo")

	categorySavedInDb := category.Category{
		ID:          "foo",
		Name:        categoryInput.Name,
		Description: categoryInput.Description,
	}
	mockedStorage.EXPECT().CreateCategory(categorySavedInDb).Return(nil)

	// WHEN
	testServer.Engine.ServeHTTP(recorder, req)
	// THEN
	assert.Equal(t, http.StatusOK, recorder.Code)
	expectedCategory, err := json.Marshal(categorySavedInDb)
	assert.NoError(t, err, "no error should be fired when we marshall the category to JSON")
	assert.Equal(t, string(expectedCategory), recorder.Body.String())

}

func TestServer_Categories(t *testing.T) {
	// GIVEN
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockedStorage := storage.NewMockStorage(ctrl)
	mockedUniqueIDGenerator := uniqueid.NewMockGenerator(ctrl)
	testServer, err := New(Config{
		Port:              8080,
		Storage:           mockedStorage,
		UniqueIDGenerator: mockedUniqueIDGenerator,
	})
	req, err := http.NewRequest("GET", "/categories", nil)
	assert.NoError(t, err, "no error should happen when building the request")
	req.Header.Add("Authorization", "ABC")
	recorder := httptest.NewRecorder()
	// mocks expectations
	mockedDBResponse := []category.Category{
		{
			ID:          "42",
			Name:        "Test",
			Description: "test",
		},
	}
	mockedStorage.EXPECT().Categories().Return(mockedDBResponse, nil)
	// WHEN
	testServer.Engine.ServeHTTP(recorder, req)
	// THEN
	assert.Equal(t, http.StatusOK, recorder.Code)
	expectedBody, err := json.Marshal(mockedDBResponse)
	assert.NoError(t, err, "no error should happen when marshalling slice of categories")
	assert.Equal(t, expectedBody, recorder.Body.Bytes())
}
