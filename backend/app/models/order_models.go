package models

import (
	"time"
)

type Order struct {
	ShippingInfo  AddressInfo `json:"shippingInfo" bson:"shippingInfo"`
	OrderItems    []*Items    `json:"orderItems" bson:"orderItems"`
	User          User        `json:"user" bson:"user"`
	PaymentInfo   Payment     `json:"paymentInfo" bson:"paymentInfo"`
	PaidAt        time.Time   `json:"paidAt" bson:"paidAt"`
	ItemsPrice    int         `json:"itemsPrice" bson:"itemsPrice"`
	TaxPrice      int         `json:"taxPrice" bson:"taxPrice"`
	ShippingPrice int         `json:"shippingPrice" bson:"shippingPrice"`
	TotalPrice    int         `json:"totalPrice" bson:"totalPrice"`
	OrderStatus   string      `json:"orderStatus" bson:"orderStatus"`
	DeliveredAt   time.Time   `json:"deliveredAt" bson:"deliveredAt"`
	CreatedAt     time.Time   `json:"createdAt" bson:"createdAt"`
}

type AddressInfo struct {
	Address string `json:"address" bson:"address"`
	City    string `json:"city" bson:"city"`
	State   string `json:"state" bson:"state"`
	Country string `json:"country" bson:"country"`
	PinCode int    `json:"pinCode" bson:"pinCode"`
	PhoneNo int    `json:"phoneNo" bson:"phoneNo"`
}

type Items struct {
	Name     string  `json:"name" bson:"name"`
	Price    int     `json:"price" bson:"price"`
	Quantity int     `json:"quantity" bson:"quantity"`
	Image    string  `json:"image" bson:"image"`
	Product  Product `json:"product" bson:"product"`
}

type Payment struct {
	Id     string `json:"id" bson:"id"`
	Status string `json:"status" bson:"status"`
}
