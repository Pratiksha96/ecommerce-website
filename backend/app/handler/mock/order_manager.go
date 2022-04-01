package mock

import (
	"ecommerce-website/app/manager"
	models "ecommerce-website/app/models"
	"testing"

	"github.com/stretchr/testify/mock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

type MockOrderManager struct {
	mock.Mock
}

// NewMockOrderManager creates a new mock instance.
func NewMockOrderManager(t *testing.T) *MockOrderManager {
	t.Helper()
	return &MockOrderManager{}
}

// CreateOrder mocks base method.
func (m *MockOrderManager) CreateOrder(order models.Order, role string, email string) (*models.Order, error) {
	args := m.Called(order, role, email)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Order), nil
}

// GetUserOrders mocks base method.
func (m *MockOrderManager) GetUserOrders(role string, email string) (manager.GetUserOrdersResponse, error) {
	args := m.Called(role, email)
	if args.Error(1) != nil {
		return manager.GetUserOrdersResponse{}, args.Error(1)
	}
	return args.Get(0).(manager.GetUserOrdersResponse), nil
}

// GetSingleOrder mocks base method.
func (m *MockOrderManager) GetSingleOrder(id primitive.ObjectID, email string) (*models.Order, error) {
	args := m.Called(id, email)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Order), nil
}

// GetAllOrders mocks base method.
func (m *MockOrderManager) GetAllOrders(role string, email string) (manager.GetAllOrdersResponse, error) {
	args := m.Called(role, email)
	if args.Error(1) != nil {
		return manager.GetAllOrdersResponse{}, args.Error(1)
	}
	return args.Get(0).(manager.GetAllOrdersResponse), nil
}

// DeleteOrder mocks base method.
func (m *MockOrderManager) DeleteOrder(id primitive.ObjectID, role string, email string) (map[string]interface{}, error) {
	args := m.Called(id, role, email)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(map[string]interface{}), nil
}

// UpdateOrder mocks base method.
func (m *MockOrderManager) UpdateOrder(status string, id primitive.ObjectID, role string, email string) (map[string]interface{}, error) {
	args := m.Called(status, id, role, email)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(map[string]interface{}), nil
}