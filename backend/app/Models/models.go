package models

type Product struct {
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
	Price       int       `json:"price" bson:"price"`
	Ratings     int       `json:"ratings" bson:"ratings"`
	Images      []*Image  `json:"images" bson:"images"`
	Category    string    `json:"category" bson:"category"`
	Stock       int       `json:"stock" bson:"stock"`
	Reviews     []*Review `json:"reviews" bson:"reviews"`
}

type Review struct {
	Name    string `json:"name" bson:"name"`
	Rating  int    `json:"rating" bson:"rating"`
	Comment string `json:"comment" bson:"comment"`
}

type Image struct {
	Public_id string `json:"public_id" bson:"public_id"`
	Url       string `json:"url" bson:"url"`
}
