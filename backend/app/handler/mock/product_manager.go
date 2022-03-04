package mock

import (
	models "ecommerce-website/app/models"
	"testing"

	"github.com/stretchr/testify/mock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

type MockProductManager struct {
	mock.Mock
}

// NewMockProductManager creates a new mock instance.
func NewMockProductManager(t *testing.T) *MockProductManager {
	t.Helper()
	return &MockProductManager{}
}

// GetProduct mocks base method.
func (m *MockProductManager) GetProduct(id primitive.ObjectID, email string) (*models.Product, error) {
	args := m.Called(id, email)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Product), nil
}

// CreateProduct mocks base method.
func (m *MockProductManager) CreateProduct(product models.Product, role string, email string) (*models.Product, error) {
	args := m.Called(product, role, email)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Product), nil
}
