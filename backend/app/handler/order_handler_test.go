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

func Test_CreateOrder(t *testing.T) {
	t.Run("When invalid request is received, it should return error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleErr := errors.New("Received invalid json request!")
		expectedReponse := utils.ErrorResponse{
			ErrorMessage: sampleErr.Error(),
			Success:      false,
		}
		expectedResponseBody, err := json.Marshal(expectedReponse)
		require.NoError(t, err)

		req, err := http.NewRequest("POST", "/order/create", strings.NewReader("{}"))
		require.NoError(t, err)

		orderManager := mock.NewMockOrderManager(t)
		handler := handler.CreateOrder(orderManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, expectedResponseBody, recorder.Body.Bytes())
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("When new order is not created, it should return error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleEmail := "sample@email.com"
		sampleOrder := utils.GetSampleOrder(sampleEmail)
		requestBody, err := json.Marshal(sampleOrder)
		require.NoError(t, err)

		sampleErr := errors.New("some error")
		expectedResponse := utils.ErrorResponse{
			ErrorMessage: sampleErr.Error(),
			Success:      false,
		}
		expectedResponseBody, err := json.Marshal(expectedResponse)
		require.NoError(t, err)

		req, err := http.NewRequest("POST", "/order/create", strings.NewReader(string(requestBody)))
		require.NoError(t, err)
		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		orderManager := mock.NewMockOrderManager(t)
		orderManager.On("CreateOrder", *sampleOrder, "user", sampleEmail).Return(nil, sampleErr)
		handler := handler.CreateOrder(orderManager)
		handler.ServeHTTP(recorder, req)

		assert.Equal(t, expectedResponseBody, recorder.Body.Bytes())
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("When order is created, it should return order successfully", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleEmail := "sample@email.com"
		sampleOrder := utils.GetSampleOrder(sampleEmail)
		requestBody, err := json.Marshal(sampleOrder)
		require.NoError(t, err)
		expectedResponse := append(requestBody, byte('\n'))

		req, err := http.NewRequest("POST", "/order/create", strings.NewReader(string(requestBody)))
		require.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		orderManager := mock.NewMockOrderManager(t)
		orderManager.On("CreateOrder", *sampleOrder, "user", sampleEmail).Return(sampleOrder, nil)

		handler := handler.CreateOrder(orderManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, string(expectedResponse), string(recorder.Body.Bytes()))
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func Test_GetUserOrders(t *testing.T) {
	t.Run("When order manager returns error, it should return error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleEmail := "sample@email.com"
		req, err := http.NewRequest("GET", "/order/user/get", nil)
		require.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))
		orderManager := mock.NewMockOrderManager(t)
		sampleErr := errors.New("some error")
		orderManager.On("GetUserOrders", "user", sampleEmail).Maybe().Return(nil, sampleErr)
		handler := handler.GetUserOrders(orderManager)
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

	t.Run("When order manager returns order, it should return order successfully", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleEmail := "sample@email.com"
		sampleOrder := &models.Order{
			User: "sample",
		}
		var doc *bson.D
		bsonReq, err := bson.Marshal(sampleOrder)
		require.NoError(t, err)
		err = bson.Unmarshal(bsonReq, &doc)
		require.NoError(t, err)
		var reqBody []primitive.M
		reqBody = append(reqBody, doc.Map())
		expectedResponse, err := json.Marshal(reqBody)
		require.NoError(t, err)
		expectedResponse = append(expectedResponse, byte('\n'))

		req, err := http.NewRequest("GET", "/order/user/get", nil)
		require.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		orderManager := mock.NewMockOrderManager(t)
		orderManager.On("GetUserOrders", "user", sampleEmail).Return(reqBody, nil)

		handler := handler.GetUserOrders(orderManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, string(expectedResponse), string(recorder.Body.Bytes()))
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func Test_GetSingleOrder(t *testing.T) {
	t.Run("When order id is not present, it should return error", func(t *testing.T) {
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
		req, err := http.NewRequest("GET", "/order/get", nil)
		require.NoError(t, err)

		vars := map[string]string{
			"id": sampleId.Hex(),
		}
		req = mux.SetURLVars(req, vars)
		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		orderManager := mock.NewMockOrderManager(t)
		sampleErr := errors.New("some error")
		orderManager.On("GetSingleOrder", sampleId, sampleEmail).Return(nil, sampleErr)
		handler := handler.GetSingleOrder(orderManager)
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

	t.Run("When order manager returns order, it should return order successfully", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleId := primitive.NewObjectID()
		sampleEmail := "sample@email.com"
		sampleOrder := &models.Order{
			User: "sample",
		}
		expectedResponse, err := json.Marshal(sampleOrder)
		require.NoError(t, err)
		expectedResponse = append(expectedResponse, byte('\n'))

		req, err := http.NewRequest("GET", "/order/get", nil)
		require.NoError(t, err)

		vars := map[string]string{
			"id": sampleId.Hex(),
		}
		req = mux.SetURLVars(req, vars)
		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		orderManager := mock.NewMockOrderManager(t)
		orderManager.On("GetSingleOrder", sampleId, sampleEmail).Return(sampleOrder, nil)

		handler := handler.GetSingleOrder(orderManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, string(expectedResponse), string(recorder.Body.Bytes()))
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func Test_GetAllOrders(t *testing.T) {
	t.Run("When no order exists, it should return error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleEmail := "sample@email.com"

		req, err := http.NewRequest("GET", "/order/get", nil)
		require.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))
		orderManager := mock.NewMockOrderManager(t)

		sampleErr := errors.New("invalid object id")
		orderManager.On("GetAllOrders", "admin", sampleEmail).Return(nil, sampleErr)
		handler := handler.GetAllOrders(orderManager)
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

	t.Run("When order exists, it should return all orders successfully", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleEmail := "sample@email.com"
		sampleOrder := &models.Order{
			User: "sample",
		}
		var doc *bson.D
		bsonReq, err := bson.Marshal(sampleOrder)
		require.NoError(t, err)
		err = bson.Unmarshal(bsonReq, &doc)
		require.NoError(t, err)
		var allOrders []primitive.M
		allOrders = append(allOrders, doc.Map())
		allOrdersResponse := manager.GetAllOrdersResponse{
			Results:     allOrders,
			TotalAmount: 99999,
		}
		expectedResponse, err := json.Marshal(allOrdersResponse)
		require.NoError(t, err)
		expectedResponse = append(expectedResponse, byte('\n'))

		req, err := http.NewRequest("GET", "/order/get", nil)
		require.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		orderManager := mock.NewMockOrderManager(t)
		orderManager.On("GetAllOrders", "admin", sampleEmail).Return(allOrdersResponse, nil)

		handler := handler.GetAllOrders(orderManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, string(expectedResponse), string(recorder.Body.Bytes()))
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func Test_DeleteOrder(t *testing.T) {
	t.Run("When order does not exist, it should return error", func(t *testing.T) {
		sampleEmail := "sample@email.com"
		recorder := httptest.NewRecorder()

		req, err := http.NewRequest("POST", "/order/delete", strings.NewReader("{}"))
		require.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		orderManager := mock.NewMockOrderManager(t)
		handler := handler.DeleteOrder(orderManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
	})

	t.Run("When unable to delete order, it should return error", func(t *testing.T) {
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

		sampleOrder := &models.Order{
			User: "sample",
		}
		requestBody, err := json.Marshal(sampleOrder)
		require.NoError(t, err)

		vars := map[string]string{
			"id": sampleId.Hex(),
		}

		req, err := http.NewRequest("POST", "/order/delete", strings.NewReader(string(requestBody)))
		require.NoError(t, err)

		req = mux.SetURLVars(req, vars)
		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		orderManager := mock.NewMockOrderManager(t)
		orderManager.On("DeleteOrder", sampleId, "admin", sampleEmail).Return(nil, sampleErr)
		handler := handler.DeleteOrder(orderManager)
		handler.ServeHTTP(recorder, req)

		assert.Equal(t, expectedResponseBody, recorder.Body.Bytes())
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("When order is deleted, it should return delete response successfully", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleId := primitive.NewObjectID()
		sampleEmail := "sample@email.com"
		sampleOrder := &models.Order{
			User: "sample",
		}
		requestBody, err := json.Marshal(sampleOrder)
		require.NoError(t, err)
		reqBody := map[string]interface{}{"success": true, "message": "document has been successfully deleted"}
		expectedResponse, err := json.Marshal(reqBody)
		require.NoError(t, err)
		expectedResponse = append(expectedResponse, byte('\n'))
		vars := map[string]string{
			"id": sampleId.Hex(),
		}

		req, err := http.NewRequest("POST", "/order/delete", strings.NewReader(string(requestBody)))
		require.NoError(t, err)

		req = mux.SetURLVars(req, vars)
		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		orderManager := mock.NewMockOrderManager(t)
		orderManager.On("DeleteOrder", sampleId, "admin", sampleEmail).Return(reqBody, nil)
		handler := handler.DeleteOrder(orderManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, string(expectedResponse), string(recorder.Body.Bytes()))
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func Test_UpdateOrder(t *testing.T) {
	t.Run("When order id is incorrect, it should return error", func(t *testing.T) {
		sampleEmail := "sample@email.com"
		recorder := httptest.NewRecorder()

		req, err := http.NewRequest("POST", "/order/update", strings.NewReader("{}"))
		require.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		orderManager := mock.NewMockOrderManager(t)
		handler := handler.DeleteOrder(orderManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
	})

	t.Run("When unable to update order, it should return error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleId := primitive.NewObjectID()
		sampleEmail := "sample@email.com"
		sampleStatus := ""
		sampleErr := errors.New("some error")
		expectedResponse := utils.ErrorResponse{
			ErrorMessage: sampleErr.Error(),
			Success:      false,
		}
		expectedResponseBody, err := json.Marshal(expectedResponse)
		require.NoError(t, err)

		sampleOrder := &models.Order{
			User: "sample",
		}
		requestBody, err := json.Marshal(sampleOrder)
		require.NoError(t, err)

		vars := map[string]string{
			"id": sampleId.Hex(),
		}

		req, err := http.NewRequest("POST", "/order/update", strings.NewReader(string(requestBody)))
		require.NoError(t, err)

		req = mux.SetURLVars(req, vars)
		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		orderManager := mock.NewMockOrderManager(t)
		orderManager.On("UpdateOrder", sampleStatus, sampleId, "admin", sampleEmail).Return(nil, sampleErr)
		handler := handler.UpdateOrder(orderManager)
		handler.ServeHTTP(recorder, req)

		assert.Equal(t, expectedResponseBody, recorder.Body.Bytes())
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("When order is updated, it should return order successfully", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleId := primitive.NewObjectID()
		sampleEmail := "sample@email.com"
		sampleStatus := ""
		sampleOrder := &models.Order{
			User: "sample",
		}
		requestBody, err := json.Marshal(sampleOrder)
		require.NoError(t, err)
		reqBody := map[string]interface{}{"success": true, "message": "document has been successfully deleted"}
		expectedResponse, err := json.Marshal(reqBody)
		require.NoError(t, err)
		expectedResponse = append(expectedResponse, byte('\n'))
		vars := map[string]string{
			"id": sampleId.Hex(),
		}

		req, err := http.NewRequest("POST", "/product/update", strings.NewReader(string(requestBody)))
		require.NoError(t, err)

		req = mux.SetURLVars(req, vars)
		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		orderManager := mock.NewMockOrderManager(t)
		orderManager.On("UpdateOrder", sampleStatus, sampleId, "admin", sampleEmail).Return(reqBody, nil)
		handler := handler.UpdateOrder(orderManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, string(expectedResponse), string(recorder.Body.Bytes()))
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
