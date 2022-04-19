package handler_test

import (
	"ecommerce-website/app/handler"
	"ecommerce-website/app/handler/mock"
	"ecommerce-website/app/models"
	"ecommerce-website/app/utils"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
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
