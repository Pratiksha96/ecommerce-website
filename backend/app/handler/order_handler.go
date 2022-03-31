package handler

import (
	"ecommerce-website/app/manager"
	models "ecommerce-website/app/models"
	"ecommerce-website/app/utils"
	"encoding/json"
	"log"
	"net/http"
)

func CreateOrder(orderManager manager.OrderManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		var order models.Order
		_ = json.NewDecoder(r.Body).Decode(&order)

		if errors := utils.OrderValidation(order); len(errors) > 0 {
			log.Println("Received invalid json request!")
			err := map[string]interface{}{"success": false, "message": errors}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}
		ctx := r.Context()
		email := ctx.Value("email").(string)
		order.OrderStatus = "processing"
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
