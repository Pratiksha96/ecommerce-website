package manager

import (
	"context"
	models "ecommerce-website/app/models"
	"ecommerce-website/internal/database"
	"log"
)

type OrderManager interface {
	CreateOrder(order models.Order, role string, email string) (*models.Order, error)
}

type orderManager struct{}

func NewOrderManager() OrderManager {
	return &orderManager{}
}

func (om *orderManager) CreateOrder(order models.Order, role string, email string) (*models.Order, error) {
	err := authorizeUser(role, email)
	if err != nil {
		return nil, err
	}
	orderPlaced, err := database.Coll_order.InsertOne(context.TODO(), order)
	if err != nil {
		return nil, err
	}

	log.Println("Order placed!", orderPlaced.InsertedID)
	return &order, nil
}
