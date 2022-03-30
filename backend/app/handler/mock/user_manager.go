package mock

import (
	"ecommerce-website/app/manager"
	models "ecommerce-website/app/models"
	"testing"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockUserManager struct {
	mock.Mock
}

// NewMockUserManager creates a new mock instance.
func NewMockUserManager(t *testing.T) *MockUserManager {
	t.Helper()
	return &MockUserManager{}
}

// RegisterUser mocks base method.
func (m *MockUserManager) RegisterUser(user models.User) (manager.TokenResponse, error) {
	args := m.Called(user)
	if args.Error(1) != nil {
		return manager.TokenResponse{}, args.Error(1)
	}
	return args.Get(0).(manager.TokenResponse), nil
}

// LoginUser mocks base method.
func (m *MockUserManager) LoginUser(user models.User) (manager.TokenResponse, error) {
	args := m.Called(user)
	if args.Error(1) != nil {
		return manager.TokenResponse{}, args.Error(1)
	}
	return args.Get(0).(manager.TokenResponse), nil
}

// GetUserDetails mocks base method.
func (m *MockUserManager) GetUserDetails(email string) (*models.User, error) {
	args := m.Called(email)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), nil
}

// UpdatePassword mocks base method
func (m *MockUserManager) UpdatePassword(email string, body map[string]interface{}) (manager.UserResponse, error) {
	args := m.Called(email, body)
	if args.Error(1) != nil {
		return manager.UserResponse{}, args.Error(1)
	}
	return args.Get(0).(manager.UserResponse), nil
}

// UpdateProfile mocks base method
func (m *MockUserManager) UpdateProfile(email string, body map[string]interface{}) (manager.UserResponse, error) {
	args := m.Called(email, body)
	if args.Error(1) != nil {
		return manager.UserResponse{}, args.Error(1)
	}
	return args.Get(0).(manager.UserResponse), nil
}

// GetAllUsers mocks base method
func (m *MockUserManager) GetAllUsers(role string, email string) ([]primitive.M, error) {
	args := m.Called(role, email)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]primitive.M), nil
}

// GetUser mocks base method
func (m *MockUserManager) GetUser(role string, email string, id primitive.ObjectID) (*models.User, error) {
	args := m.Called(role, email, id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), nil
}
