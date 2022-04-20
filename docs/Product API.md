## Get Products

Returns all the products

* **URL**

  /product/get

* **Method:**

    `GET`

* **URL Params**

  **Required:**

    `None`

* **Data Params**

    `None`

* **Success Response:**

    * **Code:** 200 <br />
      **Content:** 
```
[
{
"name" : "Book",
    "description" : "Shakespeare Othello",
    "price"  : 1000,
    "ratings" : 8,
    "category" : "romance",
    "stock" :  5,
    "images" :[ {
        "public_id" : "shkp020",
        "url" : "url://to//book"
    }],
    "reviews" : [
        {
            "name" : "user1",
            "rating" : 5,
            "comment" : "amazing book"
        },
        {
            "name" : "user2",
            "rating" : 4,
            "comment" : "good read"
        }
    ]
}
]
```

## Get Products by ID

Returns products with a given ID

* **URL**

  /product/get/{id}

* **Method:**

    `GET`

* **URL Params**

  **Required:**

    `id` exists in database

* **Data Params**

    `None`

* **Success Response:**

    * **Code:** 200 <br />
      **Content:** 
```
{
    "name" : "Book",
    "description" : "Shakespeare Othello",
    "price"  : 1000,
    "ratings" : 8,
    "category" : "romance",
    "stock" :  5,
    "images" :[ {
        "public_id" : "shkp020",
        "url" : "url://to//book"
    }],
    "reviews" : [
        {
            "name" : "user1",
            "rating" : 5,
            "comment" : "amazing book"
        },
        {
            "name" : "user2",
            "rating" : 4,
            "comment" : "good read"
        }
    ]
}
```

## Add Product

Adds a new product to the database

* **URL**

  /product/add

* **Method:**

    `POST`

* **URL Params**

  **Required:**

    `None`

* **Data Params**

    `name`
    `description`
    `price`
    `ratings`
    `category`
    `stock`
    `images`
    `reviews`

* **Success Response:**

    * **Code:** 200 <br />
      **Content:** 
```
{
    "ID": "000000000000000000000000",
    "name": "Book",
    "description": "Shakespeare Othello",
    "price": 1000,
    "ratings": 8,
    "images": [
        {
            "public_id": "shkp020",
            "url": "url://to//book"
        }
    ],
    "category": "romance",
    "stock": 5,
    "reviews": [
        {
            "user": {
                "name": "",
                "email": "",
                "password": "",
                "avatar": {
                    "public_id": "",
                    "url": ""
                },
                "role": "",
                "resetPasswordToken": "",
                "resetPasswordExpire": "0001-01-01T00:00:00Z"
            },
            "name": "user1",
            "rating": 5,
            "comment": "amazing book"
        },
        {
            "user": {
                "name": "",
                "email": "",
                "password": "",
                "avatar": {
                    "public_id": "",
                    "url": ""
                },
                "role": "",
                "resetPasswordToken": "",
                "resetPasswordExpire": "0001-01-01T00:00:00Z"
            },
            "name": "user2",
            "rating": 4,
            "comment": "good read"
        }
    ],
    "numOfReviews": 0
}
```

* **Sample Request:**
```
{
    "name" : "Book",
    "description" : "Shakespeare Othello",
    "price"  : 1000,
    "ratings" : 8,
    "category" : "romance",
    "stock" :  5,
    "images" :[ {
        "public_id" : "shkp020",
        "url" : "url://to//book"
    }],
    "reviews" : [
        {
            "name" : "user1",
            "rating" : 5,
            "comment" : "amazing book"
        },
        {
            "name" : "user2",
            "rating" : 4,
            "comment" : "good read"
        }
    ]
}
```

## Update product

Updates product with a given ID

* **URL**

  /product/update/{id}

* **Method:**

    `PUT`

* **URL Params**

  **Required:**

    `id` existing in database
 
* **Data Params**

    `name`
    `description`
    `price`
    `ratings`
    `category`
    `stock`
    `images`
    `reviews`

* **Success Response:**

    * **Code:** 200 <br />
      **Content:** 
```
{
    "name" : "Book",
    "description" : "Shakespeare Othello",
    "price"  : 500,
    "ratings" : 8,
    "category" : "romance",
    "stock" :  5,
    "images" :[ {
        "public_id" : "shkp020",
        "url" : "url://to//book"
    }],
    "reviews" : [
        {
            "name" : "user1",
            "rating" : 5,
            "comment" : "amazing book"
        },
        {
            "name" : "user2",
            "rating" : 4,
            "comment" : "good read"
        }
    ]
}
```

* **Sample Request:**
```
{
    "name" : "Book",
    "description" : "Shakespeare Othello",
    "price"  : 500,  // update price
    "ratings" : 8,
    "category" : "romance",
    "stock" :  5,
    "images" :[ {
        "public_id" : "shkp020",
        "url" : "url://to//book"
    }],
    "reviews" : [
        {
            "name" : "user1",
            "rating" : 5,
            "comment" : "amazing book"
        },
        {
            "name" : "user2",
            "rating" : 4,
            "comment" : "good read"
        }
    ]
}
```

## Delete Product

Deletes product with a given ID

* **URL**

  /product/delete/{id}

* **Method:**

    `DELETE`

* **URL Params**

  **Required:**

    `id` existing in database

* **Data Params**

    `None`

* **Success Response:**

    * **Code:** 200 <br />
      **Content:** 
```
{
    Deleted Count : 1
}
```
