package middleware

import (
	models "ecommerce-website/app/Models"
	"net/http"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductManager interface {
	CreateProduct(product models.Product, w http.ResponseWriter)
}
