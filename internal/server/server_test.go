package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

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
	assert.Equal(t, "{\"id\":\"foo\",\"name\":\"test product\",\"image\":\"\",\"shortDescription\":\"\",\"description\":\"\",\"priceVatExcluded\":{\"money\":null,\"display\":\"\"},\"vat\":{\"money\":null,\"display\":\"\"},\"totalPrice\":{\"money\":null,\"display\":\"\"}}", recorder.Body.String())
}
