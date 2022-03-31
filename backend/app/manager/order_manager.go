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
	GetUserOrders(role string, email string) (GetUserOrdersResponse, error)
	GetSingleOrder(id primitive.ObjectID, email string) (*models.Order, error)
	GetAllOrders(role string, email string) (GetAllOrdersResponse, error)
	DeleteOrder(id primitive.ObjectID, role string, email string) (map[string]interface{}, error)
}

type GetUserOrdersResponse struct {
	Results []primitive.M `json:"results" bson:"results"`
}

type GetAllOrdersResponse struct {
	Results     []primitive.M `json:"results" bson:"results"`
	TotalAmount int64         `json:"totalamount"`
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

func (om *orderManager) GetUserOrders(role string, email string) (GetUserOrdersResponse, error) {
	err := authorizeUser(role, email)
	if err != nil {
		return GetUserOrdersResponse{}, err
	}
	filter := bson.M{"user": email}

	order_list, err := database.Coll_order.Find(context.TODO(), filter)

	var results []primitive.M
	for order_list.Next(context.Background()) {
		var result bson.M
		e := order_list.Decode(&result)
		if e != nil {
			return GetUserOrdersResponse{}, err
		}
		results = append(results, result)
	}

	if err := order_list.Err(); err != nil {
		return GetUserOrdersResponse{}, err
	} else if len(results) == 0 {
		return GetUserOrdersResponse{}, errors.New("orders list is empty")
	}

	order_list.Close(context.Background())
	var response = GetUserOrdersResponse{Results: results}

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

func (om *orderManager) GetAllOrders(role string, email string) (GetAllOrdersResponse, error) {
	err := authorizeUser(role, email)
	if err != nil {
		return GetAllOrdersResponse{}, err
	}

	order_list, err := database.Coll_order.Find(context.Background(), bson.D{{}})

	var results []primitive.M
	for order_list.Next(context.Background()) {
		var result bson.M
		e := order_list.Decode(&result)
		if e != nil {
			return GetAllOrdersResponse{}, err
		}
		results = append(results, result)
	}

	if err := order_list.Err(); err != nil {
		return GetAllOrdersResponse{}, err
	} else if len(results) == 0 {
		return GetAllOrdersResponse{}, errors.New("orders list is empty")
	}

	order_list.Close(context.Background())
	var response = GetAllOrdersResponse{
		Results: results,
		//TODO to change this
		TotalAmount: 0,
	}
	return response, nil
}

func (om *orderManager) DeleteOrder(id primitive.ObjectID, role string, email string) (map[string]interface{}, error) {
	err := authorizeUser(role, email)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": id}
	deleteResult, err := database.Coll_order.DeleteOne(context.TODO(), filter)

	if err != nil {
		return nil, err
	} else if deleteResult.DeletedCount == 0 {
		return nil, errors.New("no such document present")
	}

	deleteResponse := map[string]interface{}{"success": true, "message": "order has been successfully deleted"}
	return deleteResponse, nil
}
