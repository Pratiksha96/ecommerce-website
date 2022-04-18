package mock

import (
	models "ecommerce-website/app/models"
	"testing"

	"github.com/stretchr/testify/mock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

type MockReviewManager struct {
	mock.Mock
}

// NewMockReviewManager creates a new mock instance.
func NewMockReviewManager(t *testing.T) *MockReviewManager {
	t.Helper()
	return &MockReviewManager{}
}

// CreateReview mocks base method.
func (m *MockReviewManager) CreateReview(review models.Review, product models.Product, filterProduct primitive.M) (map[string]interface{}, error) {
	args := m.Called(review, product, filterProduct)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(map[string]interface{}), nil
}

// GetProductReviews mocks base method.
func (m *MockReviewManager) GetProductReviews(id primitive.ObjectID) ([]*models.Review, error) {
	args := m.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*models.Review), nil
}

func (m *MockReviewManager) UpdateReview(review models.Review, product models.Product, filterProduct primitive.M) (map[string]interface{}, error) {
	args := m.Called(review, product, filterProduct)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(map[string]interface{}), nil
}

func (m *MockReviewManager) DeleteReview(id primitive.ObjectID, email string) (map[string]interface{}, error) {
	args := m.Called(id, email)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(map[string]interface{}), nil
}
