package manager

import (
	"context"
	models "ecommerce-website/app/models"
	"ecommerce-website/internal/database"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderManager interface {
	CreateOrder(order models.Order, role string, email string) (*models.Order, error)
	GetUserOrders(role string, email string) (GetOrdersResponse, error)
	GetSingleOrder(id primitive.ObjectID, email string) (*models.Order, error)
}

type GetOrdersResponse struct {
	Results []primitive.M `json:"results" bson:"results"`
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

func (om *orderManager) GetUserOrders(role string, email string) (GetOrdersResponse, error) {
	err := authorizeUser(role, email)
	if err != nil {
		return GetOrdersResponse{}, err
	}
	filter := bson.M{"user": email}

	order_list, err := database.Coll_order.Find(context.TODO(), filter)

	var results []primitive.M
	for order_list.Next(context.Background()) {
		var result bson.M
		e := order_list.Decode(&result)
		if e != nil {
			return GetOrdersResponse{}, err
		}
		results = append(results, result)
	}

	if err := order_list.Err(); err != nil {
		return GetOrdersResponse{}, err
	} else if len(results) == 0 {
		return GetOrdersResponse{}, errors.New("orders list is empty")
	}

	order_list.Close(context.Background())
	var response = GetOrdersResponse{Results: results}

	return response, nil
}

func (om *orderManager) GetSingleOrder(id primitive.ObjectID, email string) (*models.Order, error) {
	order := &models.Order{}
	filter := bson.M{"_id": id}
	err := database.Coll_order.FindOne(context.TODO(), filter).Decode(order)
	if err != nil {
		return order, err
	}

	return order, nil
}
