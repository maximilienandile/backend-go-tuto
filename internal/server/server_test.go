package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/maximilienandile/backend-go-tuto/internal/product"

	"github.com/stretchr/testify/assert"
)

func TestServer_CreateProduct(t *testing.T) {
	// GIVEN
	testServer, err := New(Config{
		Port: 8080,
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

	// WHEN
	testServer.Engine.ServeHTTP(recorder, req)

	// THEN
	assert.Equal(t, http.StatusOK, recorder.Code)
}
