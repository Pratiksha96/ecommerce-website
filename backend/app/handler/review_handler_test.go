package handler_test

import (
	"context"
	"ecommerce-website/app/handler"
	"ecommerce-website/app/handler/mock"
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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_GetProductReviews(t *testing.T) {
	t.Run("When product id is not present, it should return error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		req := &http.Request{}
		reviewManager := mock.NewMockReviewManager(t)
		handler := handler.GetProductReviews(reviewManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
	})

	t.Run("When review manager returns error, it should return error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleId := primitive.NewObjectID()
		req, err := http.NewRequest("GET", "/product/getreviews", nil)
		require.NoError(t, err)

		vars := map[string]string{
			"id": sampleId.Hex(),
		}
		req = mux.SetURLVars(req, vars)

		reviewManager := mock.NewMockReviewManager(t)
		sampleErr := errors.New("invalid product id")
		reviewManager.On("GetProductReviews", sampleId).Return(nil, sampleErr)
		handler := handler.GetProductReviews(reviewManager)
		handler.ServeHTTP(recorder, req)
		expectedResponse := utils.ErrorResponse{
			ErrorMessage: sampleErr.Error(),
			Success:      false,
		}
		expectedResponseBody, err := json.Marshal(expectedResponse)
		require.NoError(t, err)
		assert.Equal(t, expectedResponseBody, recorder.Body.Bytes())
		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
	})

	t.Run("When review manager return review, it should return review successfully", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleId := primitive.NewObjectID()
		sampleReview := []*models.Review{
			{
				Name: "sample review",
			},
		}
		expectedResponse, err := json.Marshal(sampleReview)
		require.NoError(t, err)
		expectedResponse = append(expectedResponse, byte('\n'))

		req, err := http.NewRequest("GET", "/product/getreviews", nil)
		require.NoError(t, err)

		vars := map[string]string{
			"id": sampleId.Hex(),
		}
		req = mux.SetURLVars(req, vars)

		reviewManager := mock.NewMockReviewManager(t)
		reviewManager.On("GetProductReviews", sampleId).Return(sampleReview, nil)

		handler := handler.GetProductReviews(reviewManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, string(expectedResponse), string(recorder.Body.Bytes()))
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func Test_DeleteReview(t *testing.T) {
	t.Run("When review is nil, it should return error", func(t *testing.T) {
		sampleEmail := "sample@email.com"
		recorder := httptest.NewRecorder()

		req, err := http.NewRequest("DELETE", "/product/deleteReview", strings.NewReader("{}"))
		require.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		reviewManager := mock.NewMockReviewManager(t)
		handler := handler.DeleteReview(reviewManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
	})

	t.Run("When unable to delete review, it should return error", func(t *testing.T) {
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

		sampleReview := []*models.Review{
			{
				Name: "sample review",
			},
		}
		requestBody, err := json.Marshal(sampleReview)
		require.NoError(t, err)

		vars := map[string]string{
			"id": sampleId.Hex(),
		}

		req, err := http.NewRequest("DELETE", "/product/deleteReview", strings.NewReader(string(requestBody)))
		require.NoError(t, err)

		req = mux.SetURLVars(req, vars)
		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		reviewManager := mock.NewMockReviewManager(t)
		reviewManager.On("DeleteReview", sampleId, sampleEmail).Return(nil, sampleErr)
		handler := handler.DeleteReview(reviewManager)
		handler.ServeHTTP(recorder, req)

		assert.Equal(t, expectedResponseBody, recorder.Body.Bytes())
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("When review is deleted, it should return delete response successfully", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleId := primitive.NewObjectID()
		sampleEmail := "sample@email.com"
		sampleReview := []*models.Review{
			{
				Name: "sample review",
			},
		}
		requestBody, err := json.Marshal(sampleReview)
		require.NoError(t, err)
		reqBody := map[string]interface{}{"success": true, "message": "document has been successfully deleted"}
		expectedResponse, err := json.Marshal(reqBody)
		require.NoError(t, err)
		expectedResponse = append(expectedResponse, byte('\n'))
		vars := map[string]string{
			"id": sampleId.Hex(),
		}

		req, err := http.NewRequest("DELETE", "/product/deleteReview", strings.NewReader(string(requestBody)))
		require.NoError(t, err)

		req = mux.SetURLVars(req, vars)
		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		reviewManager := mock.NewMockReviewManager(t)
		reviewManager.On("DeleteReview", sampleId, sampleEmail).Return(reqBody, nil)
		handler := handler.DeleteReview(reviewManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, string(expectedResponse), string(recorder.Body.Bytes()))
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
