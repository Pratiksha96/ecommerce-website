package middleware

import (
	models "ecommerce-website/app/Models"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductManager interface {
	CreateProduct(product models.Product, w http.ResponseWriter)
	GetProduct(id primitive.ObjectID, w http.ResponseWriter)
	GetAllProducts(w http.ResponseWriter)
	SearchProducts(w http.ResponseWriter, name string)
	UpdateProduct(id primitive.ObjectID, product models.Product, w http.ResponseWriter)
	DeleteProduct(id primitive.ObjectID, w http.ResponseWriter)
}
