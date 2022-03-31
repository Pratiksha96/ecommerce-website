package handler

import (
	"ecommerce-website/app/handler/mock"
	"ecommerce-website/app/manager"
	models "ecommerce-website/app/models"
	"ecommerce-website/app/utils"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_RegisterUser(t *testing.T) {
	t.Run("When invalid request is received, it should return error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		req, err := http.NewRequest("POST", "/register", strings.NewReader("{}"))
		require.NoError(t, err)

		userManager := mock.NewMockUserManager(t)
		handler := RegisterUser(userManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("When unable to register user, it should return error", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		sampleUser := &models.User{
			Name:     "sample",
			Email:    "sampleemail@email.com",
			Password: "samplepass",
			Avatar: models.ProfileImage{
				Public_id: "sampleid",
				Url:       "sampleurl",
			},
			Role:                "samplerole",
			ResetPasswordToken:  "sampletoken",
			ResetPasswordExpire: time.Now().Round(0),
		}
		requestBody, err := json.Marshal(sampleUser)
		require.NoError(t, err)

		sampleErr := errors.New("some error")
		expectedResponse := utils.ErrorResponse{
			ErrorMessage: sampleErr.Error(),
			Success:      false,
		}
		expectedResponseBody, err := json.Marshal(expectedResponse)
		require.NoError(t, err)

		req, err := http.NewRequest("POST", "register", strings.NewReader(string(requestBody)))
		require.NoError(t, err)

		userManager := mock.NewMockUserManager(t)
		userManager.On("RegisterUser", *sampleUser).Return(manager.TokenResponse{}, sampleErr)
		handler := RegisterUser(userManager)
		handler.ServeHTTP(recorder, req)

		assert.Equal(t, expectedResponseBody, recorder.Body.Bytes())
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("When able to register user, it should return response successfully", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		sampleUser := &models.User{
			Name:     "sample",
			Email:    "sampleemail@email.com",
			Password: "samplepass",
			Avatar: models.ProfileImage{
				Public_id: "sampleid",
				Url:       "sampleurl",
			},
			Role:                "samplerole",
			ResetPasswordToken:  "sampletoken",
			ResetPasswordExpire: time.Now().Round(0),
		}
		requestBody, err := json.Marshal(sampleUser)
		require.NoError(t, err)

		expectedResponse := manager.TokenResponse{
			Success: true,
			Token:   "some token",
		}
		expectedResponseBody, err := json.Marshal(expectedResponse)
		require.NoError(t, err)
		expectedResponseBody = append(expectedResponseBody, byte('\n'))

		req, err := http.NewRequest("POST", "register", strings.NewReader(string(requestBody)))
		require.NoError(t, err)

		userManager := mock.NewMockUserManager(t)
		userManager.On("RegisterUser", *sampleUser).Return(expectedResponse, nil)
		handler := RegisterUser(userManager)
		handler.ServeHTTP(recorder, req)

		assert.Equal(t, expectedResponseBody, recorder.Body.Bytes())
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func Test_LoginUser(t *testing.T) {
	t.Run("When invalid request is received, it should return error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		req, err := http.NewRequest("POST", "/login", strings.NewReader("{}"))
		require.NoError(t, err)

		userManager := mock.NewMockUserManager(t)
		handler := RegisterUser(userManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("When unable to login user, it should return error", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		sampleUser := &models.User{
			Name:     "sample",
			Email:    "sampleemail@email.com",
			Password: "samplepass",
			Avatar: models.ProfileImage{
				Public_id: "sampleid",
				Url:       "sampleurl",
			},
			Role:                "samplerole",
			ResetPasswordToken:  "sampletoken",
			ResetPasswordExpire: time.Now().Round(0),
		}
		requestBody, err := json.Marshal(sampleUser)
		require.NoError(t, err)

		sampleErr := errors.New("some error")
		expectedResponse := utils.ErrorResponse{
			ErrorMessage: sampleErr.Error(),
			Success:      false,
		}
		expectedResponseBody, err := json.Marshal(expectedResponse)
		require.NoError(t, err)

		req, err := http.NewRequest("POST", "/login", strings.NewReader(string(requestBody)))
		require.NoError(t, err)

		userManager := mock.NewMockUserManager(t)
		userManager.On("LoginUser", *sampleUser).Return(manager.TokenResponse{}, sampleErr)
		handler := LoginUser(userManager)
		handler.ServeHTTP(recorder, req)

		assert.Equal(t, expectedResponseBody, recorder.Body.Bytes())
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("When user is logged in, it should return response successfully", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		sampleUser := &models.User{
			Name:     "sample",
			Email:    "sampleemail@email.com",
			Password: "samplepass",
			Avatar: models.ProfileImage{
				Public_id: "sampleid",
				Url:       "sampleurl",
			},
			Role:                "samplerole",
			ResetPasswordToken:  "sampletoken",
			ResetPasswordExpire: time.Now().Round(0),
		}
		requestBody, err := json.Marshal(sampleUser)
		require.NoError(t, err)

		expectedResponse := manager.TokenResponse{
			Success: true,
			Token:   "some token",
		}
		expectedResponseBody, err := json.Marshal(expectedResponse)
		require.NoError(t, err)
		expectedResponseBody = append(expectedResponseBody, byte('\n'))

		req, err := http.NewRequest("POST", "/login", strings.NewReader(string(requestBody)))
		require.NoError(t, err)

		userManager := mock.NewMockUserManager(t)
		userManager.On("LoginUser", *sampleUser).Return(expectedResponse, nil)
		handler := LoginUser(userManager)
		handler.ServeHTTP(recorder, req)

		assert.Equal(t, expectedResponseBody, recorder.Body.Bytes())
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
