package handler

import (
	"context"
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

		loc, _ := time.LoadLocation("UTC")
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
			ResetPasswordExpire: time.Now().Round(0).In(loc),
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

		loc, _ := time.LoadLocation("UTC")
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
			ResetPasswordExpire: time.Now().Round(0).In(loc),
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

		loc, _ := time.LoadLocation("UTC")
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
			ResetPasswordExpire: time.Now().Round(0).In(loc),
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

		loc, _ := time.LoadLocation("UTC")
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
			ResetPasswordExpire: time.Now().Round(0).In(loc),
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

func Test_LogoutUser(t *testing.T) {
	t.Run("When cookie is not set, it should return error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		req, err := http.NewRequest("POST", "/logout", strings.NewReader("{}"))
		require.NoError(t, err)
		userManager := mock.NewMockUserManager(t)
		handler := LogoutUser(userManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	})

	t.Run("When token is set in cookie, it should delete user successfully", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		req, err := http.NewRequest("POST", "/logout", strings.NewReader("{}"))
		require.NoError(t, err)
		req.AddCookie(&http.Cookie{Name: "token", Value: "sample cookie"})

		userManager := mock.NewMockUserManager(t)
		handler := LogoutUser(userManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func Test_GetUserDetails(t *testing.T) {
	t.Run("When user does not exists, it should return error", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		sampleEmail := "sample@email.com"
		req, err := http.NewRequest("GET", "/me", nil)
		require.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		userManager := mock.NewMockUserManager(t)
		sampleErr := errors.New("some error")
		userManager.On("GetUserDetails", sampleEmail).Return(nil, sampleErr)
		handler := GetUserDetails(userManager)
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

	t.Run("When user exists, it should return user successfully", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleEmail := "sample@email.com"
		sampleUser := &models.User{
			Name: "sample user",
		}
		expectedResponse, err := json.Marshal(sampleUser)
		require.NoError(t, err)
		expectedResponse = append(expectedResponse, byte('\n'))

		req, err := http.NewRequest("GET", "/me", nil)
		require.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))

		userManager := mock.NewMockUserManager(t)
		userManager.On("GetUserDetails", sampleEmail).Return(sampleUser, nil)

		handler := GetUserDetails(userManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, string(expectedResponse), string(recorder.Body.Bytes()))
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
