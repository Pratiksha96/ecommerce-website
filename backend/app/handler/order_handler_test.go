package handler_test

import (
	"context"
	"ecommerce-website/app/handler"
	"ecommerce-website/app/handler/mock"
	"ecommerce-website/app/utils"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
