package handler_test

import (
	"bytes"
	"context"
	"ecommerce-website/app/handler"
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

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_RegisterUser(t *testing.T) {
	t.Run("When invalid request is received, it should return error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		req, err := http.NewRequest("POST", "/register", strings.NewReader("{}"))
		require.NoError(t, err)

		userManager := mock.NewMockUserManager(t)
		handler := handler.RegisterUser(userManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("When unable to register user, it should return error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleUser := utils.GetSampleUser()
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
		handler := handler.RegisterUser(userManager)
		handler.ServeHTTP(recorder, req)

		assert.Equal(t, expectedResponseBody, recorder.Body.Bytes())
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("When able to register user, it should return response successfully", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleUser := utils.GetSampleUser()
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
		handler := handler.RegisterUser(userManager)
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
		handler := handler.RegisterUser(userManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("When unable to login user, it should return error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleUser := utils.GetSampleUser()
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
		handler := handler.LoginUser(userManager)
		handler.ServeHTTP(recorder, req)

		assert.Equal(t, expectedResponseBody, recorder.Body.Bytes())
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("When user is logged in, it should return response successfully", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleUser := utils.GetSampleUser()
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
		handler := handler.LoginUser(userManager)
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
		handler := handler.LogoutUser(userManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	})

	t.Run("When token is set in cookie, it should delete user successfully", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		req, err := http.NewRequest("POST", "/logout", strings.NewReader("{}"))
		require.NoError(t, err)
		req.AddCookie(&http.Cookie{Name: "token", Value: "sample cookie"})

		userManager := mock.NewMockUserManager(t)
		handler := handler.LogoutUser(userManager)
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
		handler := handler.GetUserDetails(userManager)
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

		handler := handler.GetUserDetails(userManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, string(expectedResponse), recorder.Body.String())
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func Test_UpdatePassword(t *testing.T) {
	t.Run("When password update fails, it should return error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleEmail := "sample@email.com"
		sampleErr := errors.New("some error")
		reqBody, reqBodyBytes := map[string]interface{}{}, new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(reqBody)
		expectedResponse := utils.ErrorResponse{
			ErrorMessage: sampleErr.Error(),
			Success:      false,
		}
		expectedResponseBody, err := json.Marshal(expectedResponse)
		require.NoError(t, err)

		req, err := http.NewRequest("PUT", "/password/update", reqBodyBytes)
		require.NoError(t, err)
		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))
		userManager := mock.NewMockUserManager(t)
		userManager.On("UpdatePassword", sampleEmail, reqBody).Return(nil, sampleErr)
		handler := handler.UpdatePassword(userManager)
		handler.ServeHTTP(recorder, req)

		assert.Equal(t, expectedResponseBody, recorder.Body.Bytes())
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("When password is updated, it should return successfully", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleEmail := "sample@email.com"
		reqBody, reqBodyBytes := map[string]interface{}{"oldPassword": "old password", "newPassword": "new password",
			"confirmPassword": "confirm password"}, new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(reqBody)

		sampleUser := utils.GetSampleUser()
		expectedResponse := manager.UserResponse{
			Success: true,
			User:    *sampleUser,
		}
		expectedResponseBody, err := json.Marshal(expectedResponse)
		require.NoError(t, err)
		expectedResponseBody = append(expectedResponseBody, byte('\n'))

		req, err := http.NewRequest("PUT", "/password/update", reqBodyBytes)
		require.NoError(t, err)
		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))
		userManager := mock.NewMockUserManager(t)
		userManager.On("UpdatePassword", sampleEmail, reqBody).Return(expectedResponse, nil)
		handler := handler.UpdatePassword(userManager)
		handler.ServeHTTP(recorder, req)

		assert.Equal(t, expectedResponseBody, recorder.Body.Bytes())
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func Test_UpdateProfile(t *testing.T) {
	t.Run("When profile update fails, it should return error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleEmail := "sample@email.com"
		sampleErr := errors.New("some error")
		reqBody, reqBodyBytes := map[string]interface{}{}, new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(reqBody)
		expectedResponse := utils.ErrorResponse{
			ErrorMessage: sampleErr.Error(),
			Success:      false,
		}
		expectedResponseBody, err := json.Marshal(expectedResponse)
		require.NoError(t, err)

		req, err := http.NewRequest("PUT", "/me/update", reqBodyBytes)
		require.NoError(t, err)
		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))
		userManager := mock.NewMockUserManager(t)
		userManager.On("UpdateProfile", sampleEmail, reqBody).Return(nil, sampleErr)
		handler := handler.UpdateProfile(userManager)
		handler.ServeHTTP(recorder, req)

		assert.Equal(t, expectedResponseBody, recorder.Body.Bytes())
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("When profile is updated, it should return successfully", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleEmail := "sample@email.com"
		reqBody, reqBodyBytes := map[string]interface{}{"name": "new name", "email": "new email"}, new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(reqBody)

		sampleUser := utils.GetSampleUser()
		expectedResponse := manager.UserResponse{
			Success: true,
			User:    *sampleUser,
		}
		expectedResponseBody, err := json.Marshal(expectedResponse)
		require.NoError(t, err)
		expectedResponseBody = append(expectedResponseBody, byte('\n'))

		req, err := http.NewRequest("PUT", "/me/update", reqBodyBytes)
		require.NoError(t, err)
		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))
		userManager := mock.NewMockUserManager(t)
		userManager.On("UpdateProfile", sampleEmail, reqBody).Return(expectedResponse, nil)
		handler := handler.UpdateProfile(userManager)
		handler.ServeHTTP(recorder, req)

		assert.Equal(t, expectedResponseBody, recorder.Body.Bytes())
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func Test_GetAllUsers(t *testing.T) {
	t.Run("When user is not authorized, it should return error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleEmail := "sample@email.com"
		sampleRole := "some role"

		req, err := http.NewRequest("GET", "/user/get", nil)
		require.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))
		req = req.WithContext(context.WithValue(req.Context(), "role", sampleRole))
		userManager := mock.NewMockUserManager(t)

		sampleErr := errors.New("some error")
		userManager.On("AuthorizeUser", sampleRole, sampleEmail).Return(nil, sampleErr)
		handler := handler.GetAllUsers(userManager)
		handler.ServeHTTP(recorder, req)
		expectedResponse := utils.ErrorResponse{
			ErrorMessage: sampleErr.Error(),
			Success:      false,
		}
		expectedResponseBody, err := json.Marshal(expectedResponse)
		require.NoError(t, err)
		assert.Equal(t, string(expectedResponseBody), recorder.Body.String())
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("When user manager returns error, it should return error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleEmail := "sample@email.com"
		sampleRole := "admin"
		sampleErr := errors.New("some error")
		expectedResponse := utils.ErrorResponse{
			ErrorMessage: sampleErr.Error(),
			Success:      false,
		}
		expectedResponseBody, err := json.Marshal(expectedResponse)
		req, err := http.NewRequest("GET", "/user/get", nil)
		require.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))
		req = req.WithContext(context.WithValue(req.Context(), "role", sampleRole))
		userManager := mock.NewMockUserManager(t)

		userManager.On("AuthorizeUser", sampleRole, sampleEmail).Return(nil, sampleErr)
		userManager.On("GetAllUsers").Return(nil, sampleErr)
		handler := handler.GetAllUsers(userManager)
		handler.ServeHTTP(recorder, req)
		require.NoError(t, err)
		assert.Equal(t, string(expectedResponseBody), recorder.Body.String())
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("When user is returned successfully, it should return all users successfully", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleEmail := "sample@email.com"
		sampleRole := "admin"
		sampleUser := utils.GetSampleUser()
		var doc *bson.D
		bsonReq, err := bson.Marshal(sampleUser)
		require.NoError(t, err)
		err = bson.Unmarshal(bsonReq, &doc)
		require.NoError(t, err)
		var reqBody []primitive.M
		reqBody = append(reqBody, doc.Map())
		expectedResponse, err := json.Marshal(reqBody)
		require.NoError(t, err)
		expectedResponse = append(expectedResponse, byte('\n'))

		req, err := http.NewRequest("POST", "/user/get", strings.NewReader(string(bsonReq)))
		require.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))
		req = req.WithContext(context.WithValue(req.Context(), "role", sampleRole))

		userManager := mock.NewMockUserManager(t)
		userManager.On("AuthorizeUser", sampleRole, sampleEmail).Return(nil, nil)
		userManager.On("GetAllUsers", sampleRole, sampleEmail).Return(reqBody, nil)

		handler := handler.GetAllUsers(userManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, string(expectedResponse), recorder.Body.String())
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func Test_GetUser(t *testing.T) {
	t.Run("When user id is not present, it should return error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		req := &http.Request{}
		userManager := mock.NewMockUserManager(t)
		handler := handler.GetUser(userManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
	})

	t.Run("When user is not authorized, it should return error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleEmail := "sample@email.com"
		sampleRole := "some role"
		sampleId := primitive.NewObjectID()

		req, err := http.NewRequest("GET", "/user/get", nil)
		require.NoError(t, err)

		vars := map[string]string{
			"id": sampleId.Hex(),
		}
		req = mux.SetURLVars(req, vars)
		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))
		req = req.WithContext(context.WithValue(req.Context(), "role", sampleRole))
		userManager := mock.NewMockUserManager(t)

		sampleErr := errors.New("some error")
		userManager.On("AuthorizeUser", sampleRole, sampleEmail).Return(nil, sampleErr)
		handler := handler.GetUser(userManager)
		handler.ServeHTTP(recorder, req)
		expectedResponse := utils.ErrorResponse{
			ErrorMessage: sampleErr.Error(),
			Success:      false,
		}
		expectedResponseBody, err := json.Marshal(expectedResponse)
		require.NoError(t, err)
		assert.Equal(t, string(expectedResponseBody), recorder.Body.String())
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("When user manager returns error, it should return error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleEmail := "sample@email.com"
		sampleRole := "admin"
		sampleId := primitive.NewObjectID()
		sampleErr := errors.New("some error")
		expectedResponse := utils.ErrorResponse{
			ErrorMessage: sampleErr.Error(),
			Success:      false,
		}
		expectedResponseBody, err := json.Marshal(expectedResponse)
		req, err := http.NewRequest("GET", "/user/get", nil)
		require.NoError(t, err)

		vars := map[string]string{
			"id": sampleId.Hex(),
		}
		req = mux.SetURLVars(req, vars)
		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))
		req = req.WithContext(context.WithValue(req.Context(), "role", sampleRole))
		userManager := mock.NewMockUserManager(t)

		userManager.On("AuthorizeUser", sampleRole, sampleEmail).Return(nil, sampleErr)
		userManager.On("GetUser").Return(nil, sampleErr)
		handler := handler.GetUser(userManager)
		handler.ServeHTTP(recorder, req)
		require.NoError(t, err)
		assert.Equal(t, string(expectedResponseBody), recorder.Body.String())
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("When user is returned successfully, it should return user successfully", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		sampleEmail := "sample@email.com"
		sampleRole := "admin"
		sampleId := primitive.NewObjectID()
		sampleUser := utils.GetSampleUser()
		expectedResponse, err := json.Marshal(sampleUser)
		require.NoError(t, err)
		expectedResponse = append(expectedResponse, byte('\n'))

		req, err := http.NewRequest("GET", "/user/get", nil)
		require.NoError(t, err)

		vars := map[string]string{
			"id": sampleId.Hex(),
		}
		req = mux.SetURLVars(req, vars)
		req = req.WithContext(context.WithValue(req.Context(), "email", sampleEmail))
		req = req.WithContext(context.WithValue(req.Context(), "role", sampleRole))

		userManager := mock.NewMockUserManager(t)
		userManager.On("AuthorizeUser", sampleRole, sampleEmail).Return(nil, nil)
		userManager.On("GetUser", sampleRole, sampleEmail, sampleId).Return(sampleUser, nil)

		handler := handler.GetUser(userManager)
		handler.ServeHTTP(recorder, req)
		assert.Equal(t, string(expectedResponse), recorder.Body.String())
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
