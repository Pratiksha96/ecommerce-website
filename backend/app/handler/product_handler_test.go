package handler

import (
	"context"
	"ecommerce-website/app/handler/mock"
	"ecommerce-website/app/models"
	"ecommerce-website/app/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_GetProduct(t *testing.T) {
	t.Run("When product id is not present, it should return error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		req := &http.Request{}
		productManager := mock.NewMockProductManager(t)
		handler := GetProduct(productManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
	})

	t.Run("When product manager returns error, it should return error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleId := primitive.NewObjectID()
		sampleEmail := "sample@email.com"
		req, err := http.NewRequest("GET", "/product/get", nil)
		require.NoError(t, err)

		vars := map[string]string{
			"id": sampleId.Hex(),
		}
		req = mux.SetURLVars(req, vars)
		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		productManager := mock.NewMockProductManager(t)
		sampleErr := fmt.Errorf("some error")
		productManager.On("GetProduct", sampleId, sampleEmail).Return(nil, sampleErr)
		handler := GetProduct(productManager)
		handler.ServeHTTP(recorder, req)
		expectedResponse := utils.ErrorResponse{
			ErrorMessage: sampleErr.Error(),
			Success:      false,
		}
		expectedResponseBody, err := json.Marshal(expectedResponse)
		require.NoError(t, err)
		assert.Equal(t, expectedResponseBody, recorder.Body.Bytes())
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("When product manager return product, it should return product successfully", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleId := primitive.NewObjectID()
		sampleEmail := "sample@email.com"
		sampleProduct := &models.Product{
			Name: "sample",
		}
		expectedResponse, err := json.Marshal(sampleProduct)
		require.NoError(t, err)
		expectedResponse = append(expectedResponse, byte('\n'))

		req, err := http.NewRequest("GET", "/product/get", nil)
		require.NoError(t, err)

		vars := map[string]string{
			"id": sampleId.Hex(),
		}
		req = mux.SetURLVars(req, vars)
		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		productManager := mock.NewMockProductManager(t)
		productManager.On("GetProduct", sampleId, sampleEmail).Return(sampleProduct, nil)

		handler := GetProduct(productManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, string(expectedResponse), string(recorder.Body.Bytes()))
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func Test_CreateProduct(t *testing.T) {
	t.Run("When request is incorrect, it should return error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		req, err := http.NewRequest("POST", "/product/add", strings.NewReader("{}"))
		require.NoError(t, err)

		productManager := mock.NewMockProductManager(t)
		handler := CreateProduct(productManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("When unable to create product, it should return error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleEmail := "sample@email.com"
		sampleErr := fmt.Errorf("some error")
		expectedResponse := utils.ErrorResponse{
			ErrorMessage: sampleErr.Error(),
			Success:      false,
		}
		expectedResponseBody, err := json.Marshal(expectedResponse)
		require.NoError(t, err)

		sampleProduct := &models.Product{
			Name:        "sample",
			Description: "sample",
			Price:       700,
			Ratings:     8,
			Images: []*models.Image{
				{
					Public_id: "sampleid",
					Url:       "sampleurl",
				},
			},
			Category: "sample",
			Stock:    10,
			Reviews: []*models.Review{
				{
					Name:    "sample",
					Rating:  6,
					Comment: "sample",
				},
			},
		}
		requestBody, err := json.Marshal(sampleProduct)
		require.NoError(t, err)

		req, err := http.NewRequest("POST", "/product/add", strings.NewReader(string(requestBody)))
		require.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		productManager := mock.NewMockProductManager(t)
		productManager.On("CreateProduct", *sampleProduct, "admin", sampleEmail).Return(nil, sampleErr)
		handler := CreateProduct(productManager)
		handler.ServeHTTP(recorder, req)

		assert.Equal(t, expectedResponseBody, recorder.Body.Bytes())
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("When product is created, it should return product successfully", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleEmail := "sample@email.com"
		sampleProduct := &models.Product{
			Name:        "sample",
			Description: "sample",
			Price:       700,
			Ratings:     8,
			Images: []*models.Image{
				{
					Public_id: "sampleid",
					Url:       "sampleurl",
				},
			},
			Category: "sample",
			Stock:    10,
			Reviews: []*models.Review{
				{
					Name:    "sample",
					Rating:  6,
					Comment: "sample",
				},
			},
		}
		requestBody, err := json.Marshal(sampleProduct)
		require.NoError(t, err)
		expectedResponse := append(requestBody, byte('\n'))

		req, err := http.NewRequest("POST", "/product/add", strings.NewReader(string(requestBody)))
		require.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		productManager := mock.NewMockProductManager(t)
		productManager.On("CreateProduct", *sampleProduct, "admin", sampleEmail).Return(sampleProduct, nil)

		handler := CreateProduct(productManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, string(expectedResponse), string(recorder.Body.Bytes()))
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
