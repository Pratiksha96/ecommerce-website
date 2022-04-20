package handler

import (
	"ecommerce-website/app/manager"
	models "ecommerce-website/app/models"
	"ecommerce-website/app/utils"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateOrderRequest struct {
	Status string `json:"status"`
}

func CreateOrder(orderManager manager.OrderManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		var order models.Order
		_ = json.NewDecoder(r.Body).Decode(&order)

		if err := utils.OrderValidation(order); len(err) > 0 {
			utils.GetErrorWithStatus(errors.New("Received invalid json request!"), w, http.StatusBadRequest)
			return
		}
		ctx := r.Context()
		email := ctx.Value("email").(string)
		order.OrderStatus = "Processing"
		//adding user id for now
		order.User = email
		response, err := orderManager.CreateOrder(order, "user", email)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		json.NewEncoder(w).Encode(response)
	}
}

func GetUserOrders(orderManager manager.OrderManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		ctx := r.Context()
		email := ctx.Value("email").(string)
		orders, err := orderManager.GetUserOrders("user", email)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		json.NewEncoder(w).Encode(orders)
	}
}

func GetSingleOrder(orderManager manager.OrderManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var params = mux.Vars(r)
		id, err := primitive.ObjectIDFromHex(params["id"])
		if err != nil {
			utils.GetErrorWithStatus(errors.New("invalid order id"), w, http.StatusUnprocessableEntity)
			return
		}
		ctx := r.Context()
		email := ctx.Value("email").(string)
		order, err := orderManager.GetSingleOrder(id, email)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		json.NewEncoder(w).Encode(order)
	}
}

//for admin use only
func GetAllOrders(orderManager manager.OrderManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		ctx := r.Context()
		email := ctx.Value("email").(string)

		orders, err := orderManager.GetAllOrders("admin", email)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		json.NewEncoder(w).Encode(orders)
	}
}

//for admin use only
func DeleteOrder(orderManager manager.OrderManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var params = mux.Vars(r)
		id, err := primitive.ObjectIDFromHex(params["id"])
		if err != nil || params["id"] == "" {
			utils.GetErrorWithStatus(errors.New("invalid object id"), w, http.StatusUnprocessableEntity)
			return
		}
		ctx := r.Context()
		email := ctx.Value("email").(string)
		deleteResponse, err := orderManager.DeleteOrder(id, "admin", email)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		json.NewEncoder(w).Encode(deleteResponse)
	}
}

//for admin use only
func UpdateOrder(orderManager manager.OrderManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var params = mux.Vars(r)
		id, err := primitive.ObjectIDFromHex(params["id"])
		if err != nil || params["id"] == "" {
			utils.GetErrorWithStatus(errors.New("invalid object id"), w, http.StatusUnprocessableEntity)
			return
		}
		ctx := r.Context()
		email := ctx.Value("email").(string)
		var status UpdateOrderRequest
		_ = json.NewDecoder(r.Body).Decode(&status)

		updateResponse, err := orderManager.UpdateOrder(status.Status, id, "admin", email)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		json.NewEncoder(w).Encode(updateResponse)
	}
}
