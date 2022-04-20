## Register User

Registers a user 

* **URL**

  /me

* **Method:**

    `GET`

* **URL Params**

  **Required:**

    `None`

* **Data Params**

    `name`
    `email`
    `password`
    `avatar`

* **Success Response:**

    * **Code:** 200 <br />
      **Content:** 
```
{
    "success": true,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InN3YXRpMjJAZ21haWwuY29tIiwiZXhwIjoxNjQ2NDI3ODY2fQ.x1x1HsLnYmme8XbmTjx0ozscGFwSDhqj49hDrPMdlpo"
}
```

* **Sample Request:**
```
{
    "name" : "Test2",
	"email" : "test2@gmail.com",
	"password"  : "testUser1",
    "avatar" :{
        "public_id" : "id2",
        "url" : "url2"
    }
}
```

## Login User

Logs in a user 

* **URL**

  /login

* **Method:**

    `POST`

* **URL Params**

  **Required:**

    `None`

* **Data Params**

    `email`
    `password`

* **Success Response:**

    * **Code:** 200 <br />
      **Content:** 
```
{
    "success": true,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InN3YXRpMjJAZ21haWwuY29tIiwiZXhwIjoxNjQ2NDI3OTE0fQ.gouIqZu47jgWhsB8Nti-ZBkeTmJXfOdvA9sccARet5M"
}
```

* **Sample Request:**
```
{
	"email" : "test2@gmail.com",
    "password"  : "testUser1"
}
```

## Logs out a User

Logs a user out from application 

* **URL**

  /logout

* **Method:**

    `POST`

* **URL Params**

  **Required:**

    `None`

* **Data Params**

    `email`
    `password`

* **Success Response:**

    * **Code:** 200 <br />
      **Content:** 
```
Removes token from the cookies for the respective user
```

## Get logged-in User

Returns profile of the logged in user 

* **URL**

  /me

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

## Update Password

Updates password of user

* **URL**

  /password/update

* **Method:**

    `PUT`

* **URL Params**

  **Required:**

    `None`

* **Data Params**

    `oldPassword`
    `newPassword`
    `confirmPassowrd`

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

## Update user profile

Updates user profile details

* **URL**

  /me/update

* **Method:**

    `PUT`

* **URL Params**

  **Required:**

    `None`

* **Data Params**

    `name`
    `email`

* **Success Response:**

    * **Code:** 200 <br />
      **Content:** 
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

## Get Users

Returns all users to the admin 

* **URL**

  /user/get

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

## Get User by ID

Returns profile of a user with given ID to the admin

* **URL**

  /user/get/{id}

* **Method:**

    `GET`

* **URL Params**

  **Required:**

    `id`

* **Data Params**

    `None`

* **Success Response:**

    * **Code:** 200 <br />
      **Content:** 
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
