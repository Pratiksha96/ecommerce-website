## Create Order

Creates an order

* **URL**

  /order/create

* **Method:**

    `POST`

* **URL Params**

  **Required:**

    `None`

* **Data Params**

    `shippingInfo`
    `orderItems`
    `itemsPrice`
    `taxPrice`
    `shippingPrice`
    `totalPrice`
    `paymentInfo`

* **Sample Request:**
```
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
```

* **Success Response:**

    * **Code:** 200 <br />
      **Content:** 
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


## Get all user specific Orders

Returns all previous orders of a user specific to it's email id i.e. identified by user logged-in id

* **URL**

  /order/user/get

* **Method:**

    `GET`

* **URL Params**

  **Required:**

    `None`

* **Data Params**

    `None`

* **Sample Request:**
```
http://localhost:8081/order/user/get
```

* **Success Response:**

    * **Code:** 200 <br />
      **Content:** 
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
        }
    ]
}
]
```


## Get one Order

Returns a specific order based on its id

* **URL**

  /order/get/{id}

* **Method:**

    `GET`

* **URL Params**

  **Required:**

    `None`

* **Data Params**

    `None`

* **Sample Request:**
```
http://localhost:8081/order/get/624693e666527ec1a1f9e00b
```

* **Success Response:**

    * **Code:** 200 <br />
      **Content:** 
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

## Get all existing orders

Returns all orders for an admin to see

* **URL**

  /order/get

* **Method:**

    `GET`

* **URL Params**

  **Required:**

    `None`

* **Data Params**

    `None`

* **Sample Request:**
```
http://localhost:8081/order/get
```

* **Success Response:**

    * **Code:** 200 <br />
      **Content:** 
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
        }
    ],
    "totalamount": 0
}
]
```

## Delete an Order

Returns success or failed response after deleting an order

* **URL**

  /order/delete/{id}

* **Method:**

    `DELETE`

* **URL Params**

  **Required:**

    `id` exists in database

* **Data Params**

    `None`

* **Sample Request:**
```
http://localhost:8081/order/delete/62454ee3ae51f7bfde45f798 or http://localhost:8081/order/delete/624693e666527ec1a1f9e00b
```

* **Success Response:**

    * **Code:** 200 <br />
      **Content:** 
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


## Update an Order

Returns an updated order by updating the delivery status of the respective order. 
At the same time, it will reduce the product count for all products present in that order.

* **URL**

  /order/update/{id}

* **Method:**

    `PUT`

* **URL Params**

  **Required:**

    `id` order id exists in database

* **Data Params**

    `status` as Shipped

* **Sample Request:**
```
http://localhost:8081/order/update/624693e666527ec1a1f9e00c
```

* **Success Response:**

    * **Code:** 200 <br />
      **Content:** 
```
[
    {
    "message": "Password updated succesfully",
    "success": "true"
}
]
```
