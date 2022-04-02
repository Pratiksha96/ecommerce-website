# Sprint 2

## Stories
Developer:
- [Implementation of User manager should be wrapped under an interface](https://github.com/Pratiksha96/ecommerce-website/issues/99)
- [As a developer, I want to add validations for order related details](https://github.com/Pratiksha96/ecommerce-website/issues/131)
- [As a developer, I want to add Order model to start storing customer orders](https://github.com/Pratiksha96/ecommerce-website/issues/130)


Admin:
- [As an admin, I want to update an order placed by a user](https://github.com/Pratiksha96/ecommerce-website/issues/141)
- [As an admin, I want to be able to delete an order](https://github.com/Pratiksha96/ecommerce-website/issues/140)
- [As an admin, I want to see all orders placed by others users](https://github.com/Pratiksha96/ecommerce-website/issues/139)
- [As an admin, I want to fetch details of all existing users](https://github.com/Pratiksha96/ecommerce-website/issues/103)
- [As an admin, I want to fetch details of a specific user](https://github.com/Pratiksha96/ecommerce-website/issues/104)


User:
- [As a user, I want to see all my orders placed](https://github.com/Pratiksha96/ecommerce-website/issues/136)
- [As a User, I want to get all my details after I login](https://github.com/Pratiksha96/ecommerce-website/issues/100)
- [As a User, I want to be able to update my profile](https://github.com/Pratiksha96/ecommerce-website/issues/101)
- [As a user, I want to be able to create orders](https://github.com/Pratiksha96/ecommerce-website/issues/133)
- [As a user, I want to be able to see a specific order](https://github.com/Pratiksha96/ecommerce-website/issues/138)
- [As a user, I want to add a product review](https://github.com/Pratiksha96/ecommerce-website/issues/110)



Unit Tests:
- [Unit test for Get All Products Handler function](https://github.com/Pratiksha96/ecommerce-website/issues/45)
- [Unit test for Delete Product Handler function](https://github.com/Pratiksha96/ecommerce-website/issues/76)
- [Unit test for Search Products Handler function](https://github.com/Pratiksha96/ecommerce-website/issues/73)
- [Unit test for Register User Handler function](https://github.com/Pratiksha96/ecommerce-website/issues/129)
- [Unit test for Login User Handler function](https://github.com/Pratiksha96/ecommerce-website/issues/128)
- [Unit test for Logout User Handler function](https://github.com/Pratiksha96/ecommerce-website/issues/127)
- [Unit test for Get User Details Handler function](https://github.com/Pratiksha96/ecommerce-website/issues/126)
- [Unit test for Update Password Handler function](https://github.com/Pratiksha96/ecommerce-website/issues/125)
- [Unit test for Update Profile Handler function](https://github.com/Pratiksha96/ecommerce-website/issues/124)
- [Unit tests for Get User Handler function](https://github.com/Pratiksha96/ecommerce-website/issues/45)
- [Unit tests for Get All Users Handler function](https://github.com/Pratiksha96/ecommerce-website/issues/147)
- [Unit test for Create Order handler function](https://github.com/Pratiksha96/ecommerce-website/issues/153)

Bug:
- [Sending incorrect product count on search product query](https://github.com/Pratiksha96/ecommerce-website/issues/114)

# Backend

## Database Schema
- A database named "ecommerce_website" exists where we added a collection named Order.
- Schema of the Order is as -
```
   { 
    ShippingInfo  AddressInfo 
	OrderItems    []*Items    
	User          string      
	PaymentInfo   Payment     
	PaidAt        time.Time  
	ItemsPrice    int         
	TaxPrice      int
	ShippingPrice int
	TotalPrice    int
	OrderStatus   string
	DeliveredAt   time.Time
	CreatedAt     time.Time
   }
```
- AddressInfo schema is - 
```
    {
    Address string
    City    string
    State   string
    Country string
    PinCode int   
    PhoneNo int   
    }
```
- Items schema is - 
```
    {
    Name     string        
    Price    int               
    Quantity int               
    Image    string            
    Product  primitive.ObjectID
    }
```
- Payment schema is - 
```
    {
    Id     string
    Status string
    }
```
- Review in product schema is changed to - 
```
    {
	User    User   
	Name    string 
	Rating  int    
	Comment string 
    }
```
## API Contracts
The endpoints developed in this sprint for user are - 

```
/me
/password/update
/me/update
/user/get
/user/get/{id}
/product/createReview
```

The endpoints developed in this sprint for order are - 
```
/order/create
/order/user/get
/order/get/{id}
/order/get
/order/delete/{id}
/order/update/{id}
```



### GET /me

### PUT /password/update

### PUT /me/update

### PUT /product/createReview

### GET /user/get

### GET /user/get/{id}

### POST /order/create

Returns user created orders
Sample: http://localhost:8081/order/create
```
URL Params
None
Data Params
{
    "shippingInfo" : {
        "address" : "4000 SW",
        "city" : "Gainesville",
        "state" : "Florida",
        "country" : "Us",
        "pinCode" : 32608,
        "phoneNo" : 9898989898
    },
    "orderItems" : [{
        "product" : "62227595235a5df8404aac47",
        "name" : "Men's Red T-shirt",
        "price" : 10,
        "image" : "sample image",
        "quantity" : 4
    },{
        "product" : "622278da235a5df8404aac49",
        "name" : "Men's Green T-shirt",
        "price" : 15,
        "image" : "sample image",
        "quantity" : 2
    }],
	"itemsPrice" : 200,
	"taxPrice"  : 36,
	"shippingPrice" : 100,
    "totalPrice" : 336,
    "paymentInfo":{
        "id" : "sample info",
        "status" : "succeeded"
    }
}
Headers
Content-Type: application/json

Success Response:
Code: 200
Content:
```
```
[
    {
    "shippingInfo": {
        "address": "4000 SW",
        "city": "Gainesville",
        "state": "Florida",
        "country": "Us",
        "pinCode": 32608,
        "phoneNo": 9898989898
    },
    "orderItems": [
        {
            "name": "Men's Red T-shirt",
            "price": 10,
            "quantity": 4,
            "image": "sample image",
            "product": "62227595235a5df8404aac47"
        },
        {
            "name": "Men's Green T-shirt",
            "price": 15,
            "quantity": 2,
            "image": "sample image",
            "product": "622278da235a5df8404aac49"
        }
    ],
    "user": "admin2@gmail.com",
    "paymentInfo": {
        "id": "sample info",
        "status": "succeeded"
    },
    "paidAt": "0001-01-01T00:00:00Z",
    "itemsPrice": 200,
    "taxPrice": 36,
    "shippingPrice": 100,
    "totalPrice": 336,
    "orderStatus": "Processing",
    "deliveredAt": "0001-01-01T00:00:00Z",
    "createdAt": "0001-01-01T00:00:00Z"
}
]
``` 

### GET /order/user/get
Returns all previous orders of a user specific to it's email id
Sample: http://localhost:8081/order/user/get
```
URL Params
None
Data Params
None
Headers
Content-Type: application/json

Success Response:
Code: 200
Content:
```
```
[
    {
    "results": [
        {
            "_id": "624693d866527ec1a1f9e00a",
            "createdAt": "0000-12-31T19:03:58-04:56",
            "deliveredAt": "2022-04-01T02:13:15.102-04:00",
            "itemsPrice": 200,
            "orderItems": [
                {
                    "image": "sample image",
                    "name": "Men's Red T-shirt",
                    "price": 10,
                    "product": "62227595235a5df8404aac47",
                    "quantity": 4
                }
            ],
            "orderStatus": "Shipped",
            "paidAt": "0000-12-31T19:03:58-04:56",
            "paymentInfo": {
                "id": "sample info",
                "status": "succeeded"
            },
            "shippingInfo": {
                "address": "4000 SW",
                "city": "Gainesville",
                "country": "Us",
                "phoneNo": 9898989898,
                "pinCode": 32608,
                "state": "Florida"
            },
            "shippingPrice": 100,
            "taxPrice": 36,
            "totalPrice": 336,
            "user": "admin2@gmail.com"
        },
        {
            "_id": "624693e666527ec1a1f9e00b",
            "createdAt": "0000-12-31T19:03:58-04:56",
            "deliveredAt": "0000-12-31T19:03:58-04:56",
            "itemsPrice": 200,
            "orderItems": [
                {
                    "image": "sample image",
                    "name": "Men's Red T-shirt",
                    "price": 10,
                    "product": "62227595235a5df8404aac47",
                    "quantity": 4
                }
            ],
            "orderStatus": "processing",
            "paidAt": "0000-12-31T19:03:58-04:56",
            "paymentInfo": {
                "id": "sample info",
                "status": "succeeded"
            },
            "shippingInfo": {
                "address": "4000 SW",
                "city": "Gainesville",
                "country": "Us",
                "phoneNo": 9898989898,
                "pinCode": 32608,
                "state": "Florida"
            },
            "shippingPrice": 100,
            "taxPrice": 36,
            "totalPrice": 336,
            "user": "admin2@gmail.com"
        },
        {
            "_id": "624693e666527ec1a1f9e00c",
            "createdAt": "0000-12-31T19:03:58-04:56",
            "deliveredAt": "0000-12-31T19:03:58-04:56",
            "itemsPrice": 200,
            "orderItems": [
                {
                    "image": "sample image",
                    "name": "Men's Red T-shirt",
                    "price": 10,
                    "product": "62227595235a5df8404aac47",
                    "quantity": 4
                }
            ],
            "orderStatus": "processing",
            "paidAt": "0000-12-31T19:03:58-04:56",
            "paymentInfo": {
                "id": "sample info",
                "status": "succeeded"
            },
            "shippingInfo": {
                "address": "4000 SW",
                "city": "Gainesville",
                "country": "Us",
                "phoneNo": 9898989898,
                "pinCode": 32608,
                "state": "Florida"
            },
            "shippingPrice": 100,
            "taxPrice": 36,
            "totalPrice": 336,
            "user": "admin2@gmail.com"
        },
        {
            "_id": "624693e766527ec1a1f9e00d",
            "createdAt": "0000-12-31T19:03:58-04:56",
            "deliveredAt": "0000-12-31T19:03:58-04:56",
            "itemsPrice": 200,
            "orderItems": [
                {
                    "image": "sample image",
                    "name": "Men's Red T-shirt",
                    "price": 10,
                    "product": "62227595235a5df8404aac47",
                    "quantity": 4
                }
            ],
            "orderStatus": "processing",
            "paidAt": "0000-12-31T19:03:58-04:56",
            "paymentInfo": {
                "id": "sample info",
                "status": "succeeded"
            },
            "shippingInfo": {
                "address": "4000 SW",
                "city": "Gainesville",
                "country": "Us",
                "phoneNo": 9898989898,
                "pinCode": 32608,
                "state": "Florida"
            },
            "shippingPrice": 100,
            "taxPrice": 36,
            "totalPrice": 336,
            "user": "admin2@gmail.com"
        },
        {
            "_id": "6246951ea2779eec5c613bc8",
            "createdAt": "0000-12-31T19:03:58-04:56",
            "deliveredAt": "0000-12-31T19:03:58-04:56",
            "itemsPrice": 200,
            "orderItems": [
                {
                    "image": "sample image",
                    "name": "Men's Red T-shirt",
                    "price": 10,
                    "product": "62227595235a5df8404aac47",
                    "quantity": 4
                }
            ],
            "orderStatus": "Processing",
            "paidAt": "0000-12-31T19:03:58-04:56",
            "paymentInfo": {
                "id": "sample info",
                "status": "succeeded"
            },
            "shippingInfo": {
                "address": "4000 SW",
                "city": "Gainesville",
                "country": "Us",
                "phoneNo": 9898989898,
                "pinCode": 32608,
                "state": "Florida"
            },
            "shippingPrice": 100,
            "taxPrice": 36,
            "totalPrice": 336,
            "user": "admin2@gmail.com"
        },
        {
            "_id": "62469b81cc7864c611ce1622",
            "createdAt": "2022-04-01T02:28:17.762-04:00",
            "deliveredAt": "0000-12-31T19:03:58-04:56",
            "itemsPrice": 200,
            "orderItems": [
                {
                    "image": "sample image",
                    "name": "Men's Red T-shirt",
                    "price": 10,
                    "product": "62227595235a5df8404aac47",
                    "quantity": 4
                },
                {
                    "image": "sample image",
                    "name": "Men's Green T-shirt",
                    "price": 15,
                    "product": "622278da235a5df8404aac49",
                    "quantity": 2
                }
            ],
            "orderStatus": "Shipped",
            "paidAt": "2022-04-01T02:28:17.762-04:00",
            "paymentInfo": {
                "id": "sample info",
                "status": "succeeded"
            },
            "shippingInfo": {
                "address": "4000 SW",
                "city": "Gainesville",
                "country": "Us",
                "phoneNo": 9898989898,
                "pinCode": 32608,
                "state": "Florida"
            },
            "shippingPrice": 100,
            "taxPrice": 36,
            "totalPrice": 336,
            "user": "admin2@gmail.com"
        },
        {
            "_id": "62477fe1ebf833b053a04e5d",
            "createdAt": "0000-12-31T19:03:58-04:56",
            "deliveredAt": "0000-12-31T19:03:58-04:56",
            "itemsPrice": 200,
            "orderItems": [
                {
                    "image": "sample image",
                    "name": "Men's Red T-shirt",
                    "price": 10,
                    "product": "62227595235a5df8404aac47",
                    "quantity": 4
                },
                {
                    "image": "sample image",
                    "name": "Men's Green T-shirt",
                    "price": 15,
                    "product": "622278da235a5df8404aac49",
                    "quantity": 2
                }
            ],
            "orderStatus": "Processing",
            "paidAt": "0000-12-31T19:03:58-04:56",
            "paymentInfo": {
                "id": "sample info",
                "status": "succeeded"
            },
            "shippingInfo": {
                "address": "4000 SW",
                "city": "Gainesville",
                "country": "Us",
                "phoneNo": 9898989898,
                "pinCode": 32608,
                "state": "Florida"
            },
            "shippingPrice": 100,
            "taxPrice": 36,
            "totalPrice": 336,
            "user": "admin2@gmail.com"
        }
    ]
}
]
``` 

### GET /order/get/{id}
Returns a specific order
Sample: http://localhost:8081/order/get/624693e666527ec1a1f9e00b
```
URL Params
Order Id
Data Params
None
Headers
Content-Type: application/json

Success Response:
Code: 200
Content:
```
```
[
    {
    "shippingInfo": {
        "address": "4000 SW",
        "city": "Gainesville",
        "state": "Florida",
        "country": "Us",
        "pinCode": 32608,
        "phoneNo": 9898989898
    },
    "orderItems": [
        {
            "name": "Men's Red T-shirt",
            "price": 10,
            "quantity": 4,
            "image": "sample image",
            "product": "62227595235a5df8404aac47"
        }
    ],
    "user": "admin2@gmail.com",
    "paymentInfo": {
        "id": "sample info",
        "status": "succeeded"
    },
    "paidAt": "0001-01-01T00:00:00Z",
    "itemsPrice": 200,
    "taxPrice": 36,
    "shippingPrice": 100,
    "totalPrice": 336,
    "orderStatus": "processing",
    "deliveredAt": "0001-01-01T00:00:00Z",
    "createdAt": "0001-01-01T00:00:00Z"
}
]
``` 

### GET /order/get
Returns all orders for an admin to see
Sample: http://localhost:8081/order/get
```
URL Params
None
Data Params
None
Headers
Content-Type: application/json

Success Response:
Code: 200
Content:
```
```
[
    {
    "results": [
        {
            "_id": "624693d866527ec1a1f9e00a",
            "createdAt": "0000-12-31T19:03:58-04:56",
            "deliveredAt": "2022-04-01T02:13:15.102-04:00",
            "itemsPrice": 200,
            "orderItems": [
                {
                    "image": "sample image",
                    "name": "Men's Red T-shirt",
                    "price": 10,
                    "product": "62227595235a5df8404aac47",
                    "quantity": 4
                }
            ],
            "orderStatus": "Shipped",
            "paidAt": "0000-12-31T19:03:58-04:56",
            "paymentInfo": {
                "id": "sample info",
                "status": "succeeded"
            },
            "shippingInfo": {
                "address": "4000 SW",
                "city": "Gainesville",
                "country": "Us",
                "phoneNo": 9898989898,
                "pinCode": 32608,
                "state": "Florida"
            },
            "shippingPrice": 100,
            "taxPrice": 36,
            "totalPrice": 336,
            "user": "admin2@gmail.com"
        },
        {
            "_id": "624693e666527ec1a1f9e00b",
            "createdAt": "0000-12-31T19:03:58-04:56",
            "deliveredAt": "0000-12-31T19:03:58-04:56",
            "itemsPrice": 200,
            "orderItems": [
                {
                    "image": "sample image",
                    "name": "Men's Red T-shirt",
                    "price": 10,
                    "product": "62227595235a5df8404aac47",
                    "quantity": 4
                }
            ],
            "orderStatus": "processing",
            "paidAt": "0000-12-31T19:03:58-04:56",
            "paymentInfo": {
                "id": "sample info",
                "status": "succeeded"
            },
            "shippingInfo": {
                "address": "4000 SW",
                "city": "Gainesville",
                "country": "Us",
                "phoneNo": 9898989898,
                "pinCode": 32608,
                "state": "Florida"
            },
            "shippingPrice": 100,
            "taxPrice": 36,
            "totalPrice": 336,
            "user": "admin2@gmail.com"
        },
        {
            "_id": "624693e666527ec1a1f9e00c",
            "createdAt": "0000-12-31T19:03:58-04:56",
            "deliveredAt": "0000-12-31T19:03:58-04:56",
            "itemsPrice": 200,
            "orderItems": [
                {
                    "image": "sample image",
                    "name": "Men's Red T-shirt",
                    "price": 10,
                    "product": "62227595235a5df8404aac47",
                    "quantity": 4
                }
            ],
            "orderStatus": "processing",
            "paidAt": "0000-12-31T19:03:58-04:56",
            "paymentInfo": {
                "id": "sample info",
                "status": "succeeded"
            },
            "shippingInfo": {
                "address": "4000 SW",
                "city": "Gainesville",
                "country": "Us",
                "phoneNo": 9898989898,
                "pinCode": 32608,
                "state": "Florida"
            },
            "shippingPrice": 100,
            "taxPrice": 36,
            "totalPrice": 336,
            "user": "admin2@gmail.com"
        },
        {
            "_id": "624693e766527ec1a1f9e00d",
            "createdAt": "0000-12-31T19:03:58-04:56",
            "deliveredAt": "0000-12-31T19:03:58-04:56",
            "itemsPrice": 200,
            "orderItems": [
                {
                    "image": "sample image",
                    "name": "Men's Red T-shirt",
                    "price": 10,
                    "product": "62227595235a5df8404aac47",
                    "quantity": 4
                }
            ],
            "orderStatus": "processing",
            "paidAt": "0000-12-31T19:03:58-04:56",
            "paymentInfo": {
                "id": "sample info",
                "status": "succeeded"
            },
            "shippingInfo": {
                "address": "4000 SW",
                "city": "Gainesville",
                "country": "Us",
                "phoneNo": 9898989898,
                "pinCode": 32608,
                "state": "Florida"
            },
            "shippingPrice": 100,
            "taxPrice": 36,
            "totalPrice": 336,
            "user": "admin2@gmail.com"
        },
        {
            "_id": "6246951ea2779eec5c613bc8",
            "createdAt": "0000-12-31T19:03:58-04:56",
            "deliveredAt": "0000-12-31T19:03:58-04:56",
            "itemsPrice": 200,
            "orderItems": [
                {
                    "image": "sample image",
                    "name": "Men's Red T-shirt",
                    "price": 10,
                    "product": "62227595235a5df8404aac47",
                    "quantity": 4
                }
            ],
            "orderStatus": "Processing",
            "paidAt": "0000-12-31T19:03:58-04:56",
            "paymentInfo": {
                "id": "sample info",
                "status": "succeeded"
            },
            "shippingInfo": {
                "address": "4000 SW",
                "city": "Gainesville",
                "country": "Us",
                "phoneNo": 9898989898,
                "pinCode": 32608,
                "state": "Florida"
            },
            "shippingPrice": 100,
            "taxPrice": 36,
            "totalPrice": 336,
            "user": "admin2@gmail.com"
        },
        {
            "_id": "62469b81cc7864c611ce1622",
            "createdAt": "2022-04-01T02:28:17.762-04:00",
            "deliveredAt": "0000-12-31T19:03:58-04:56",
            "itemsPrice": 200,
            "orderItems": [
                {
                    "image": "sample image",
                    "name": "Men's Red T-shirt",
                    "price": 10,
                    "product": "62227595235a5df8404aac47",
                    "quantity": 4
                },
                {
                    "image": "sample image",
                    "name": "Men's Green T-shirt",
                    "price": 15,
                    "product": "622278da235a5df8404aac49",
                    "quantity": 2
                }
            ],
            "orderStatus": "Shipped",
            "paidAt": "2022-04-01T02:28:17.762-04:00",
            "paymentInfo": {
                "id": "sample info",
                "status": "succeeded"
            },
            "shippingInfo": {
                "address": "4000 SW",
                "city": "Gainesville",
                "country": "Us",
                "phoneNo": 9898989898,
                "pinCode": 32608,
                "state": "Florida"
            },
            "shippingPrice": 100,
            "taxPrice": 36,
            "totalPrice": 336,
            "user": "admin2@gmail.com"
        },
        {
            "_id": "62477fe1ebf833b053a04e5d",
            "createdAt": "0000-12-31T19:03:58-04:56",
            "deliveredAt": "0000-12-31T19:03:58-04:56",
            "itemsPrice": 200,
            "orderItems": [
                {
                    "image": "sample image",
                    "name": "Men's Red T-shirt",
                    "price": 10,
                    "product": "62227595235a5df8404aac47",
                    "quantity": 4
                },
                {
                    "image": "sample image",
                    "name": "Men's Green T-shirt",
                    "price": 15,
                    "product": "622278da235a5df8404aac49",
                    "quantity": 2
                }
            ],
            "orderStatus": "Processing",
            "paidAt": "0000-12-31T19:03:58-04:56",
            "paymentInfo": {
                "id": "sample info",
                "status": "succeeded"
            },
            "shippingInfo": {
                "address": "4000 SW",
                "city": "Gainesville",
                "country": "Us",
                "phoneNo": 9898989898,
                "pinCode": 32608,
                "state": "Florida"
            },
            "shippingPrice": 100,
            "taxPrice": 36,
            "totalPrice": 336,
            "user": "admin2@gmail.com"
        }
    ],
    "totalamount": 0
}
]
``` 

### DELETE /order/delete/{id}
Returns success or failed response after deleting an order document
Sample: http://localhost:8081/order/delete/62454ee3ae51f7bfde45f798 or http://localhost:8081/order/delete/624693e666527ec1a1f9e00b

```
URL Params
Order Id
Data Params
None
Headers
Content-Type: application/json

Success Response:
Code: 200
Content:
```
```
[
    {
    "success": false,
    "message": "no such document present"
    }
]

or

[
    {
    "message": "order has been successfully deleted",
    "success": true
    }
]

``` 

### PUT /order/update/{id}
Returns an updated order by updating the delivery status of the respective order. 
At the same time, it will reduce the product count for all products present in that order.
Sample: http://localhost:8081/order/update/624693e666527ec1a1f9e00c
```
URL Params
Order Id
Data Params
{
    "status" : "Shipped"
}
Headers
Content-Type: application/json

Success Response:
Code: 200
Content:
```
```
[
    {
    "message": "order has been updated successfully",
    "success": true
}
]
```

### GET /me
Returns profile of the logged in user 
Sample: http://localhost:8081/me

```
URL Params
None
Data Params
None
Headers
Content-Type: application/json

Success Response:
Code: 200
Content:
```
```
[
    {
    "name": "test22",
    "email": "test22@gmail.com",
    "password": "$2a$10$8b0vCuZ4Wn493BTWWlvwgOHnJOB0sTYgsrlJckvt2vfRv.bDZmG2e",
    "avatar": {
        "public_id": "test_id1",
        "url": "test_url1"
    },
    "role": "user",
    "resetPasswordToken": "",
    "resetPasswordExpire": "0001-01-01T00:00:00Z"
}
]
```

### PUT /passowrd/update
User can update passoword using this link
Sample: http://localhost:8081/password/update

```
URL Params
None
Data Params
{
	"oldPassword" : "**********",
    "newPassword"  : "**********",
    "confirmPassowrd" : "**********"
}
Headers
Content-Type: application/json

Success Response:
Code: 200
Content:
```
```
[
    {
    "message": "Password updated succesfully",
    "success": "true"
}
]
```

### PUT /me/update
User can update his profile details.
Sample: http://localhost:8081/me/update

```
URL Params
None
Data Params
{
    "name" : "test22",
	"email" : "test22@gmail.com"
}
Headers
Content-Type: application/json

Success Response:
Code: 200
Content:
```
```
[
    {
    "Success": true,
    "User": {
        "name": "test22",
        "email": "test22@gmail.com",
        "password": "$2a$10$8b0vCuZ4Wn493BTWWlvwgOHnJOB0sTYgsrlJckvt2vfRv.bDZmG2e",
        "avatar": {
            "public_id": "test_id1",
            "url": "test_url1"
        },
        "role": "user",
        "resetPasswordToken": "",
        "resetPasswordExpire": "0001-01-01T00:00:00Z"
    }
}
]
```

### GET /user/get
Admin can get details of all the users using this link
Sample: http://localhost:8081/getAllUsers

```
URL Params
None
Data Params
None
Headers
Content-Type: application/json

Success Response:
Code: 200
Content:
```
```
[
    {
    "Success": true,
    "User": {
        "name": "test22",
        "email": "test22@gmail.com",
        "password": "$2a$10$8b0vCuZ4Wn493BTWWlvwgOHnJOB0sTYgsrlJckvt2vfRv.bDZmG2e",
        "avatar": {
            "public_id": "test_id1",
            "url": "test_url1"
        },
        "role": "user",
        "resetPasswordToken": "",
        "resetPasswordExpire": "0001-01-01T00:00:00Z"
    }
}
]
```

### GET /user/get/{id}
Admin can get details of all the users using this link
Sample: http://localhost:8081/user/get/622278a4235a5df8404aac44

```
URL Params
{id}
Data Params
None
Headers
Content-Type: application/json

Success Response:
Code: 200
Content:
```
```
[
    {
    "Success": true,
    "User": {
        "name": "test22",
        "email": "test22@gmail.com",
        "password": "$2a$10$8b0vCuZ4Wn493BTWWlvwgOHnJOB0sTYgsrlJckvt2vfRv.bDZmG2e",
        "avatar": {
            "public_id": "test_id1",
            "url": "test_url1"
        },
        "role": "user",
        "resetPasswordToken": "",
        "resetPasswordExpire": "0001-01-01T00:00:00Z"
    }
}
]
```

### POST product/createReview
User can get details of a particular order he placed
Sample: http://localhost:8081/product/createReview

```
URL Params
None
Data Params
{
    "productId" : "62227595235a5df8404aac47",
    "rating" : "5",
    "comment" : "Average Quality Product, not so great"

}
Headers
Content-Type: application/json

Success Response:
Code: 200
Content:
```
```
[
    {
    "message": "Review has been created",
    "success": true
}
]
```

## Backend API Development
- REST apis for placing orders using Go lang added specific to user, admin and developer.
- Orders have been filtered for users after verifying the user with its email id
- Admin can fetch details of all users or particular order. Only admin can access this resource.
- User can Update his password as well as profile details.
- User can also add product reviews for a particular product. It will change the number of reviews on products as well as avg rating of product.
- All of the development was done VS code.
- The server was hosted using mux router and the APIs were configured on this.
- APIs were also configured with MongoDB database using new order collection, and were tested using POSTMAN tool.


## Test results

### Testing /order/create api

![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/testImages/CreateOrder.png)

### Testing /order/user/get api

![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/testImages/GetUserSpecificOrder.png)

### Testing /order/get/{order_id} api

![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/testImages/GetSingleOrder.png)

### Testing /order/get api

![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/testImages/GetAllOrders.png)

### Testing /order/delete/{order_id} api

![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/testImages/DeleteOrder.png)

### Testing /order/update/{order_id} api

![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/testImages/UpdateOrder.png)


### GET /me

![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/testImages/getProfile.png)

### PUT /me/update

![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/testImages/updateProfile.png)

### PUT /product/createReview

![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/testImages/createReview.png)

# Frontend 

## FrontEnd Technologies: React JS, CSS
## User Profile
A layout is created where a user can see his profile on the website. The page is mobile responsive. Here are the screenshots to  see the profile of the user for different screen sizes.
## UpdateProfile
A layout is created where a user can update his profile on the website. The page is mobile responsive. Here are the screenshots to  see the update profile of the user for different screen sizes.

## UpdatePassword
A layout is created where a user can update his password on the website. The page is mobile responsive. Here are the screenshots to  see the update password of the user for different screen sizes.
## Cart
A layout is created where a user can add products in a cart on the website. The page is mobile responsive. Here are the screenshots to  see the checkout page for different screen sizes.

##  User Menu
A layout is created where a user can see the user menu on the website. Here are the screenshots to  see the user menu.


## Test results 
###  User Menu
![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/testImages/UserMenu.png)

### User Profile
![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/testImages/UserProfile.png)
### User Profile-mobile
![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/testImages/UserProfile-mobile.png)
### UpdateProfile
![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/testImages/UpdateProfile.png)
### UpdateProfile-mobile
![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/testImages/UpdateProfile-mobile.png)
### UpdatePassword
![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/testImages/UpdatePassword.png)
### UpdatePassword-mobile
![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/testImages/UpdatePassword-mobile.png)

### Cart
![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/testImages/Cart.png)

### Cart-mobile
![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/testImages/Cart-mobile.png)

## How to run
### Backend
- download mongo db - make sure it is running by running command - `mongo`
- checkout main branch and use 'make run' for running the server
- use 'make build'
- It should show in your terminal that Db is connected.
- Verify server is running by running this - localhost:8081/ping - it should result in pong on your screen

### Frontend
 - download node - make sure it is running by running command - `node -v`
- checkout main branch and cd into frontend
- run npm install to install all the dependencies 
- run npm start to start the server
- Verify server is running by running this - localhost:3000 - it should display the home pages