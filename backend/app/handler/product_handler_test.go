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
	"go.mongodb.org/mongo-driver/bson"
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

func Test_GetAllProducts(t *testing.T) {
	t.Run("When no product exists, it should return error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleEmail := "sample@email.com"

		req, err := http.NewRequest("GET", "/product/get", nil)
		require.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))
		productManager := mock.NewMockProductManager(t)

		sampleErr := fmt.Errorf("invalid object id")
		productManager.On("GetAllProducts", sampleEmail).Return(nil, sampleErr)
		handler := GetAllProducts(productManager)
		handler.ServeHTTP(recorder, req)
		expectedResponse := utils.ErrorResponse{
			ErrorMessage: sampleErr.Error(),
			Success:      false,
		}
		expectedResponseBody, err := json.Marshal(expectedResponse)
		require.NoError(t, err)
		assert.Equal(t, string(expectedResponseBody), string(recorder.Body.Bytes()))
		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
	})

	t.Run("When products exists, it should return all products successfully", func(t *testing.T) {
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

		var doc *bson.D
		bsonReq, err := bson.Marshal(sampleProduct)
		require.NoError(t, err)
		err = bson.Unmarshal(bsonReq, &doc)
		require.NoError(t, err)
		var reqBody []primitive.M
		reqBody = append(reqBody, doc.Map())

		req, err := http.NewRequest("POST", "/product/add", strings.NewReader(string(bsonReq)))
		require.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		productManager := mock.NewMockProductManager(t)
		productManager.On("GetAllProducts", sampleEmail).Return(reqBody, nil)

		handler := GetAllProducts(productManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, reqBody, string(recorder.Body.Bytes()))
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func Test_UpdateProduct(t *testing.T) {
	t.Run("When request is incorrect, it should return error", func(t *testing.T) {
		sampleEmail := "sample@email.com"
		sampleId := primitive.NewObjectID()
		recorder := httptest.NewRecorder()

		req, err := http.NewRequest("POST", "/product/update", strings.NewReader("{}"))
		require.NoError(t, err)

		vars := map[string]string{
			"id": sampleId.Hex(),
		}
		req = mux.SetURLVars(req, vars)
		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		productManager := mock.NewMockProductManager(t)
		handler := UpdateProduct(productManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("When unable to update product, it should return error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleId := primitive.NewObjectID()
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

		vars := map[string]string{
			"id": sampleId.Hex(),
		}

		req, err := http.NewRequest("POST", "/product/add", strings.NewReader(string(requestBody)))
		require.NoError(t, err)

		req = mux.SetURLVars(req, vars)
		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		productManager := mock.NewMockProductManager(t)
		productManager.On("UpdateProduct", sampleId, *sampleProduct, "admin", sampleEmail).Return(nil, sampleErr)
		handler := UpdateProduct(productManager)
		handler.ServeHTTP(recorder, req)

		assert.Equal(t, expectedResponseBody, recorder.Body.Bytes())
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("When product is created, it should return product successfully", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleId := primitive.NewObjectID()
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

		vars := map[string]string{
			"id": sampleId.Hex(),
		}

		req, err := http.NewRequest("POST", "/product/add", strings.NewReader(string(requestBody)))
		require.NoError(t, err)

		req = mux.SetURLVars(req, vars)
		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		productManager := mock.NewMockProductManager(t)
		productManager.On("UpdateProduct", sampleId, *sampleProduct, "admin", sampleEmail).Return(sampleProduct, nil)
		handler := UpdateProduct(productManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, string(expectedResponse), string(recorder.Body.Bytes()))
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func Test_DeleteProduct(t *testing.T) {
	t.Run("When product is nil, it should return error", func(t *testing.T) {
		sampleEmail := "sample@email.com"
		recorder := httptest.NewRecorder()

		req, err := http.NewRequest("POST", "/product/delete", strings.NewReader("{}"))
		require.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		productManager := mock.NewMockProductManager(t)
		handler := DeleteProduct(productManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
	})

	t.Run("When unable to delete product, it should return error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleId := primitive.NewObjectID()
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

		vars := map[string]string{
			"id": sampleId.Hex(),
		}

		req, err := http.NewRequest("POST", "/product/add", strings.NewReader(string(requestBody)))
		require.NoError(t, err)

		req = mux.SetURLVars(req, vars)
		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		productManager := mock.NewMockProductManager(t)
		productManager.On("DeleteProduct", sampleId, "admin", sampleEmail).Return(nil, sampleErr)
		handler := DeleteProduct(productManager)
		handler.ServeHTTP(recorder, req)

		assert.Equal(t, expectedResponseBody, recorder.Body.Bytes())
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("When product is deleted, it should return delete response successfully", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleId := primitive.NewObjectID()
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
		expectedResponse := map[string]interface{}{"success": true, "message": "document has been successfully deleted"}
		vars := map[string]string{
			"id": sampleId.Hex(),
		}

		req, err := http.NewRequest("POST", "/product/add", strings.NewReader(string(requestBody)))
		require.NoError(t, err)

		req = mux.SetURLVars(req, vars)
		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		productManager := mock.NewMockProductManager(t)
		productManager.On("DeleteProduct", sampleId, "admin", sampleEmail).Return(expectedResponse, nil)
		handler := DeleteProduct(productManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, expectedResponse, string(recorder.Body.Bytes()))
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
