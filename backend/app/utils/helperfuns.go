package utils

import (
	models "ecommerce-website/app/models"
	"encoding/json"
	"log"
	"net/http"
	"net/mail"
	"net/url"
)

type ErrorResponse struct {
	Success      bool   `json:"success"`
	ErrorMessage string `json:"message"`
}

func GetError(err error, w http.ResponseWriter) {
	log.Printf(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		Success:      false,
	}
	message, _ := json.Marshal(response)

	w.WriteHeader(http.StatusInternalServerError)
	w.Write(message)
}

func GetErrorWithStatus(err error, w http.ResponseWriter, statusCode int) {
	log.Printf(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		Success:      false,
	}
	message, _ := json.Marshal(response)

	w.WriteHeader(statusCode)
	w.Write(message)
}

func ProductValidation(product models.Product) url.Values {
	errors := url.Values{}

	if product.Name == "" {
		errors.Add("Name", "Please enter product name!")
	}

	if product.Description == "" {
		errors.Add("Description", "Please enter product description!")
	}

	if product.Price <= 0 {
		errors.Add("Price", "Please enter valid product price!")
	}

	if product.Price > 99999999 {
		errors.Add("Prices", "Product price can not exceed length 8!")
	}

	if product.Ratings <= 0 {
		errors.Add("Ratings", "Product ratings can not be negative or empty!")
	}

	if len(product.Images) == 0 {
		errors.Add("Images", "Product images can not be empty!")
	}

	if product.Category == "" {
		errors.Add("Category", "Product category can not be empty!")
	}

	if product.Stock <= 0 {
		errors.Add("Stock", "Please enter product stock!")
	}

	return errors
}

func UserRegisterValidation(user models.User) url.Values {

	errors := url.Values{}

	if user.Name == "" {
		errors.Add("name", "Please enter user name!")
	}

	if len(user.Name) > 30 || len(user.Name) < 4 {
		errors.Add("name", "User name should lie between length 4 and 30!")
	}

	if len(user.Email) == 0 {
		errors.Add("email", "User email address is mandatory!")
	}

	if len(user.Password) < 8 {
		errors.Add("password", "User password should be atleast 8 characters long!")
	}

	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		errors.Add("email", "Invalid email address given!")
	}

	return errors
}

func UserLoginValidation(user models.User) url.Values {

	errors := url.Values{}

	if len(user.Email) == 0 {
		errors.Add("email", "User email address is mandatory!")
	}

	if len(user.Password) == 0 {
		errors.Add("password", "User password is required to login!")
	}

	return errors
}

//initial validations, might alter later
func OrderValidation(order models.Order) url.Values {

	errors := url.Values{}

	if (models.AddressInfo{}) == order.ShippingInfo {
		errors.Add("shippingInfo", "Please enter shipping address!")
	}
	if (models.Payment{}) == order.PaymentInfo {
		errors.Add("paymentInfo", "Please enter payment info!")
	}
	if order.TotalPrice <= 0 {
		errors.Add("totalPrice", "Total price is missing!")
	}
	if len(order.OrderItems) == 0 {
		errors.Add("orderItems", "Order items are empty!")
	}

	if order.ItemsPrice <= 0 {
		errors.Add("itemsPrice", "Item price is missing!")
	}

	if order.TaxPrice <= 0 {
		errors.Add("taxPrice", "Tax price is missing!")
	}

	if order.ShippingPrice <= 0 {
		errors.Add("shippingPrice", "Shipping price is missing!")
	}

	return errors
}
