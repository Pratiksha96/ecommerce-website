## Create Review

Create a new review for a particular product

* **URL**

  /product/createReview

* **Method:**

    `PUT`

* **URL Params**

  **Required:**

    `None`

* **Data Params**

    `productId`
    `rating`
    `comment`
    

* **Success Response:**

    * **Code:** 200 <br />
      **Content:** 
```
[
{
    "message": "Review has been created",
    "success": true
}
]
```

* **Sample Request:**
```
{
    "productId" : "62227595235a5df8404aac47",
    "rating" : "5",
    "comment" : "Very good,Excellent Quality Product"
}
```

## Get Product Reviews

Returns all reviews of a particular product

* **URL**

  /product/getreviews/{id}

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
 [
 
    {
        "user": {
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
        },
        "name": "test22",
        "rating": 4,
        "comment": "Average Quality Product, not so great"
    },
    {
        "user": {
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
        },
        "name": "test22",
        "rating": 5,
        "comment": "Average Quality Product, not so great"
    },
    {
        "user": {
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
        },
        "name": "test22",
        "rating": 5,
        "comment": "Average Quality Product, not so great"
    }
]
```

## Update Review

Updates review given by a user

* **URL**

  /product/updateReview

* **Method:**

    `PUT`

* **URL Params**

  **Required:**

    `None`

* **Data Params**

    `productId`
    `rating`
    `comment`

* **Success Response:**

    * **Code:** 200 <br />
      **Content:** 
```
{
    "message": "Review has been updated",
    "success": true
}
```

* **Sample Request:**
```
[
    {
    "productId" : "62227595235a5df8404aac47",
    "rating" : "3",
    "comment" : "Average Product"
    }
]
```

## Delete Review

Deletes Review added by a user for particular product

* **URL**

  /product/deleteReview/{id}

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
    "message": "Review has been successfully deleted",
    "success": true
}
```
