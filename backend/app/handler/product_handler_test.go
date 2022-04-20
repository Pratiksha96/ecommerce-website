package handler_test

import (
	"context"
	"ecommerce-website/app/handler"
	"ecommerce-website/app/handler/mock"
	"ecommerce-website/app/manager"
	"ecommerce-website/app/models"
	"ecommerce-website/app/utils"
	"encoding/json"
	"errors"
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
		handler := handler.GetProduct(productManager)
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
		sampleErr := errors.New("some error")
		productManager.On("GetProduct", sampleId, sampleEmail).Return(nil, sampleErr)
		handler := handler.GetProduct(productManager)
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

		handler := handler.GetProduct(productManager)
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
		handler := handler.CreateProduct(productManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("When unable to create product, it should return error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleEmail := "sample@email.com"
		sampleErr := errors.New("some error")
		expectedResponse := utils.ErrorResponse{
			ErrorMessage: sampleErr.Error(),
			Success:      false,
		}
		expectedResponseBody, err := json.Marshal(expectedResponse)
		require.NoError(t, err)

		sampleProduct := utils.GetSampleProduct()
		requestBody, err := json.Marshal(sampleProduct)
		require.NoError(t, err)

		req, err := http.NewRequest("POST", "/product/add", strings.NewReader(string(requestBody)))
		require.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		productManager := mock.NewMockProductManager(t)
		productManager.On("CreateProduct", *sampleProduct, "admin", sampleEmail).Return(nil, sampleErr)
		handler := handler.CreateProduct(productManager)
		handler.ServeHTTP(recorder, req)

		assert.Equal(t, expectedResponseBody, recorder.Body.Bytes())
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("When product is created, it should return product successfully", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleEmail := "sample@email.com"
		sampleProduct := utils.GetSampleProduct()
		requestBody, err := json.Marshal(sampleProduct)
		require.NoError(t, err)
		expectedResponse := append(requestBody, byte('\n'))

		req, err := http.NewRequest("POST", "/product/add", strings.NewReader(string(requestBody)))
		require.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		productManager := mock.NewMockProductManager(t)
		productManager.On("CreateProduct", *sampleProduct, "admin", sampleEmail).Return(sampleProduct, nil)

		handler := handler.CreateProduct(productManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, string(expectedResponse), recorder.Body.String())
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

		sampleErr := errors.New("invalid object id")
		productManager.On("GetAllProducts").Return(nil, sampleErr)
		handler := handler.GetAllProducts(productManager)
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
		sampleProduct := utils.GetSampleProduct()
		var doc *bson.D
		bsonReq, err := bson.Marshal(sampleProduct)
		require.NoError(t, err)
		err = bson.Unmarshal(bsonReq, &doc)
		require.NoError(t, err)
		var reqBody []primitive.M
		reqBody = append(reqBody, doc.Map())
		expectedResponse, err := json.Marshal(reqBody)
		require.NoError(t, err)
		expectedResponse = append(expectedResponse, byte('\n'))

		req, err := http.NewRequest("GET", "/product/get", nil)
		require.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		productManager := mock.NewMockProductManager(t)
		productManager.On("GetAllProducts").Return(reqBody, nil)

		handler := handler.GetAllProducts(productManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, string(expectedResponse), string(recorder.Body.Bytes()))
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func Test_UpdateProduct(t *testing.T) {
	t.Run("When request is incorrect, it should return error", func(t *testing.T) {
		sampleEmail := "sample@email.com"
		sampleId := primitive.NewObjectID()
		recorder := httptest.NewRecorder()

		req, err := http.NewRequest("PUT", "/product/update", strings.NewReader("{}"))
		require.NoError(t, err)

		vars := map[string]string{
			"id": sampleId.Hex(),
		}
		req = mux.SetURLVars(req, vars)
		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		productManager := mock.NewMockProductManager(t)
		handler := handler.UpdateProduct(productManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("When unable to update product, it should return error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleId := primitive.NewObjectID()
		sampleEmail := "sample@email.com"
		sampleErr := errors.New("some error")
		expectedResponse := utils.ErrorResponse{
			ErrorMessage: sampleErr.Error(),
			Success:      false,
		}
		expectedResponseBody, err := json.Marshal(expectedResponse)
		require.NoError(t, err)

		sampleProduct := utils.GetSampleProduct()
		requestBody, err := json.Marshal(sampleProduct)
		require.NoError(t, err)

		vars := map[string]string{
			"id": sampleId.Hex(),
		}

		req, err := http.NewRequest("PUT", "/product/update", strings.NewReader(string(requestBody)))
		require.NoError(t, err)

		req = mux.SetURLVars(req, vars)
		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		productManager := mock.NewMockProductManager(t)
		productManager.On("UpdateProduct", sampleId, *sampleProduct, "admin", sampleEmail).Return(nil, sampleErr)
		handler := handler.UpdateProduct(productManager)
		handler.ServeHTTP(recorder, req)

		assert.Equal(t, expectedResponseBody, recorder.Body.Bytes())
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("When product is updated, it should return product successfully", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleId := primitive.NewObjectID()
		sampleEmail := "sample@email.com"
		sampleProduct := utils.GetSampleProduct()
		requestBody, err := json.Marshal(sampleProduct)
		require.NoError(t, err)
		expectedResponse := append(requestBody, byte('\n'))

		vars := map[string]string{
			"id": sampleId.Hex(),
		}

		req, err := http.NewRequest("PUT", "/product/update", strings.NewReader(string(requestBody)))
		require.NoError(t, err)

		req = mux.SetURLVars(req, vars)
		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		productManager := mock.NewMockProductManager(t)
		productManager.On("UpdateProduct", sampleId, *sampleProduct, "admin", sampleEmail).Return(sampleProduct, nil)
		handler := handler.UpdateProduct(productManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, string(expectedResponse), string(recorder.Body.Bytes()))
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func Test_DeleteProduct(t *testing.T) {
	t.Run("When product is nil, it should return error", func(t *testing.T) {
		sampleEmail := "sample@email.com"
		recorder := httptest.NewRecorder()

		req, err := http.NewRequest("DELETE", "/product/delete", strings.NewReader("{}"))
		require.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		productManager := mock.NewMockProductManager(t)
		handler := handler.DeleteProduct(productManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
	})

	t.Run("When unable to delete product, it should return error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleId := primitive.NewObjectID()
		sampleEmail := "sample@email.com"
		sampleErr := errors.New("some error")
		expectedResponse := utils.ErrorResponse{
			ErrorMessage: sampleErr.Error(),
			Success:      false,
		}
		expectedResponseBody, err := json.Marshal(expectedResponse)
		require.NoError(t, err)

		sampleProduct := utils.GetSampleProduct()
		requestBody, err := json.Marshal(sampleProduct)
		require.NoError(t, err)

		vars := map[string]string{
			"id": sampleId.Hex(),
		}

		req, err := http.NewRequest("DELETE", "/product/delete", strings.NewReader(string(requestBody)))
		require.NoError(t, err)

		req = mux.SetURLVars(req, vars)
		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		productManager := mock.NewMockProductManager(t)
		productManager.On("DeleteProduct", sampleId, "admin", sampleEmail).Return(nil, sampleErr)
		handler := handler.DeleteProduct(productManager)
		handler.ServeHTTP(recorder, req)

		assert.Equal(t, expectedResponseBody, recorder.Body.Bytes())
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("When product is deleted, it should return delete response successfully", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleId := primitive.NewObjectID()
		sampleEmail := "sample@email.com"
		sampleProduct := utils.GetSampleProduct()
		requestBody, err := json.Marshal(sampleProduct)
		require.NoError(t, err)
		reqBody := map[string]interface{}{"success": true, "message": "document has been successfully deleted"}
		expectedResponse, err := json.Marshal(reqBody)
		require.NoError(t, err)
		expectedResponse = append(expectedResponse, byte('\n'))
		vars := map[string]string{
			"id": sampleId.Hex(),
		}

		req, err := http.NewRequest("DELETE", "/product/delete", strings.NewReader(string(requestBody)))
		require.NoError(t, err)

		req = mux.SetURLVars(req, vars)
		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		productManager := mock.NewMockProductManager(t)
		productManager.On("DeleteProduct", sampleId, "admin", sampleEmail).Return(reqBody, nil)
		handler := handler.DeleteProduct(productManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, string(expectedResponse), string(recorder.Body.Bytes()))
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func Test_SearchProducts(t *testing.T) {
	t.Run("When search manager returns error, it should return error", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		req, err := http.NewRequest("GET", "/product/search", nil)
		require.NoError(t, err)

		q := req.URL.Query()

		q.Add("keyword", "hello")
		q.Add("priceMax", "700")
		q.Add("page", "1")

		req.URL.RawQuery = q.Encode()

		productManager := mock.NewMockProductManager(t)

		sampleErr := errors.New("invalid object id")
		productManager.On("SearchProducts", q).Return(manager.SearchResponse{}, sampleErr)
		handler := handler.SearchProducts(productManager)
		handler.ServeHTTP(recorder, req)
		expectedResponse := utils.ErrorResponse{
			ErrorMessage: sampleErr.Error(),
			Success:      false,
		}
		expectedResponseBody, err := json.Marshal(expectedResponse)
		require.NoError(t, err)
		assert.Equal(t, string(expectedResponseBody), string(recorder.Body.Bytes()))
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("When search results exists, it should return searched products successfully", func(t *testing.T) {

		recorder := httptest.NewRecorder()
		sampleProduct := utils.GetSampleProduct()
		var doc *bson.D
		bsonReq, err := bson.Marshal(sampleProduct)
		require.NoError(t, err)
		err = bson.Unmarshal(bsonReq, &doc)
		require.NoError(t, err)
		var searchResult []primitive.M
		searchResult = append(searchResult, doc.Map())

		response := manager.SearchResponse{
			Results:       searchResult,
			TotalProducts: 1,
		}
		expectedResponse, err := json.Marshal(response)
		require.NoError(t, err)
		expectedResponse = append(expectedResponse, byte('\n'))

		req, err := http.NewRequest("GET", "/product/search", nil)
		require.NoError(t, err)

		q := req.URL.Query()

		q.Add("keyword", "hello")
		q.Add("priceMax", "700")
		q.Add("page", "1")

		req.URL.RawQuery = q.Encode()

		productManager := mock.NewMockProductManager(t)
		productManager.On("SearchProducts", q).Return(response, nil)

		handler := handler.SearchProducts(productManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, string(expectedResponse), string(recorder.Body.Bytes()))
		assert.Equal(t, http.StatusOK, recorder.Code)

	})
}
