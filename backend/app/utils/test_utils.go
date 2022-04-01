package utils

import (
	"ecommerce-website/app/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetSampleProduct() *models.Product {
	sampleProduct := &models.Product{
		Name:        "sample",
		Description: "sample",
		Price:       700,
		Ratings:     8,
		Images: []*models.Image{
			{
				Public_id: "sampleid",
				Url:       "sampleurl",
			},
		},
		Category: "sample",
		Stock:    10,
		Reviews: []*models.Review{
			{
				Name:    "sample",
				Rating:  6,
				Comment: "sample",
			},
		},
	}
	return sampleProduct
}

func GetSampleOrder(sampleEmail string) *models.Order {
	loc, _ := time.LoadLocation("UTC")
	sampleOrder := &models.Order{
		User: sampleEmail,
		ShippingInfo: models.AddressInfo{
			Address: "sample address",
			City:    "sample city",
			State:   "sample state",
			Country: "sample country",
			PinCode: 1234,
			PhoneNo: 5678,
		},
		PaymentInfo: models.Payment{
			Id:     "sample id",
			Status: "sample status",
		},
		TotalPrice: 10,
		OrderItems: []*models.Items{
			{
				Name:     "sample name",
				Price:    10,
				Quantity: 1,
				Image:    "some image",
				Product:  primitive.NewObjectID(),
			},
		},
		ItemsPrice:    12,
		TaxPrice:      5,
		ShippingPrice: 17,
		PaidAt:        time.Now().Round(0).In(loc),
		DeliveredAt:   time.Now().Round(0).In(loc),
		CreatedAt:     time.Now().Round(0).In(loc),
		OrderStatus:   "Processing",
	}
	return sampleOrder
}

func GetSampleUser() *models.User {
	loc, _ := time.LoadLocation("UTC")
	sampleUser := &models.User{
		Name:     "sample",
		Email:    "sampleemail@email.com",
		Password: "samplepass",
		Avatar: models.ProfileImage{
			Public_id: "sampleid",
			Url:       "sampleurl",
		},
		Role:                "samplerole",
		ResetPasswordToken:  "sampletoken",
		ResetPasswordExpire: time.Now().Round(0).In(loc),
	}
	return sampleUser
}
