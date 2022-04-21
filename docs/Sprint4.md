<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Sprint 4](#sprint-4)
  - [Stories](#stories)
- [Backend](#backend)
  - [Database Schema](#database-schema)
  - [API Contracts](#api-contracts)
    - [PUT /product/createReview](#put-productcreatereview)
    - [GET /product/getreviews/{id}](#get-productgetreviewsid)
    - [PUT /product/updateReview](#put-productupdatereview)
    - [DELETE /product/deleteReview/{id}](#delete-productdeletereviewid)
  - [Backend API Development](#backend-api-development)
  - [Test results](#test-results)
- [Frontend](#frontend)
  - [FrontEnd Technologies: React JS, CSS](#frontend-technologies-react-js-css)
  - [Test results](#test-results-1)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Sprint 4

## Stories
Developer:
- [As a developer, I want unit test for Create Order handler](https://github.com/Pratiksha96/ecommerce-website/issues/167)
- [As a developer, I want unit test for Update Order handler](https://github.com/Pratiksha96/ecommerce-website/issues/165)
- [As a developer, I want unit test for Get All Orders handler](https://github.com/Pratiksha96/ecommerce-website/issues/168)
- [As a developer, I want unit test for Get Single Order handler](https://github.com/Pratiksha96/ecommerce-website/issues/169)
- [As a developer, I want unit test for Get User Orders handler](https://github.com/Pratiksha96/ecommerce-website/issues/156)
- [As a developer, I want unit test for Create Review handler](https://github.com/Pratiksha96/ecommerce-website/issues/175)
- [As a developer, I want unit test for Update Review handler](https://github.com/Pratiksha96/ecommerce-website/issues/177)
- [As a developer, I want unit test for Delete Review handler](https://github.com/Pratiksha96/ecommerce-website/issues/178)
- [As a developer, I want unit test for Get Product Reviews handler](https://github.com/Pratiksha96/ecommerce-website/issues/176)
- [As a developer, I don't want to allow similar role update for an existing user](https://github.com/Pratiksha96/ecommerce-website/issues/189)
- [As a developer, I want to give an option of changing user/admin roles](https://github.com/Pratiksha96/ecommerce-website/issues/106)
- [As a developer, I want to test frontend.](https://github.com/Pratiksha96/ecommerce-website/issues/185)
- [As a developer, I want efficient method to calculate average rating of product.](https://github.com/Pratiksha96/ecommerce-website/issues/188)
- [As a developer, I want to give an option of showing product reviews for a product](https://github.com/Pratiksha96/ecommerce-website/issues/111)

Admin:
- [As an admin, I want to be able to delete an existing user](https://github.com/Pratiksha96/ecommerce-website/issues/107)
- [As an admin, I want to have different routes for admin and user. ](https://github.com/Pratiksha96/ecommerce-website/issues/186)

User:
- [As a user, I want to add and delete products in a cart. ](https://github.com/Pratiksha96/ecommerce-website/issues/183)
- [As a user, I want to review the product. ](https://github.com/Pratiksha96/ecommerce-website/issues/184)
- [As a user, I want to create an order. ](https://github.com/Pratiksha96/ecommerce-website/issues/187)
- [As a user, I want to be able to delete product review](https://github.com/Pratiksha96/ecommerce-website/issues/112)
- [As a user I want to be able to update my product review](https://github.com/Pratiksha96/ecommerce-website/issues/172)


Bug:
- [As a developer, I want roles other than 'user' or 'admin' to be declined in the role change request](https://github.com/Pratiksha96/ecommerce-website/issues/192)
- [As a developer, I should not accept negative product stock or product prices or order prices etc](https://github.com/Pratiksha96/ecommerce-website/issues/198)
- [As a developer I don't want users to submit negative rating or rating greater than 5](https://github.com/Pratiksha96/ecommerce-website/issues/196)



# Backend
## Database Schema
No new schema was added. Worked on existing schema involving reviews. 
- Schema of a review is - 
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
/user/updaterole
/user/delete/{id}

```
The endpoints developed in this sprint for user are - 
```
/product/getreviews/{id}
/product/updateReview
/product/deleteReview/{id}

```

### PUT /user/updaterole
Returns a success or failed response after changing the role of the requested user
Sample: http://localhost:8081/user/updaterole
```
URL Params
None
Data Params
{
    "id" : "622278a4235a5df8404aac48",
    "role" : "user"
}
Headers
Content-Type: application/json

Success Response:
Code: 200
Content:
```
```
{
    "Success": true,
    "User": {
        "name": "Test1",
        "email": "test1@gmail.com",
        "password": "$2a$10$O3Yi//j0DxMp2gqtfI/1hONSG/dceTquVqVVAtugu.DvzHOyWJ5cm",
        "avatar": {
            "public_id": "test_id1",
            "url": "test_url1"
        },
        "role": "user",
        "resetPasswordToken": "",
        "resetPasswordExpire": "0001-01-01T00:00:00Z"
    }
}

or

{
    "success": false,
    "message": "The requested role has already been assigned to this id. Hence, no change"
}

or

{
    "success": false,
    "message": "The role in the change request is invalid!"
}

``` 

### DELETE /user/delete/{id}

Returns a success or failed response after deleting the requested user
Sample: http://localhost:8081/user/delete/622291405e49b77d8dd8ced6
```
URL Params
id of an user
Data Params
None
Headers
Content-Type: application/json

Success Response:
Code: 200
Content:
```
```
{
    "message": "user has been successfully deleted",
    "success": true
}

or

{
    "success": false,
    "message": "no such user present"
}

``` 

### GET /product/getreviews/{id}
Returns list of all reviews for a particular product
Sample: http://localhost:8081/product/getreviews/62227595235a5df8404aac47
```
URL Params
id of product
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

### PUT /product/updateReview
Returns a success or failed response after updating review of a product
Sample: http://localhost:8081//product/updateReview
```
URL Params
None
Data Params
{
    "productId" : "62227595235a5df8404aac47",
    "rating" : "5",
    "comment" : "Very Good Product
}
Headers
Content-Type: application/json

Success Response:
Code: 200
Content:
```
```
{
    "message": "Review has been updated",
    "success": true
}
```

### DELETE /product/deleteReview/{id}

Returns a success or failed response after deleting review of a particular user for given product
Sample: http://localhost:8081/product/deleteReview/62227595235a5df8404aac47
```
URL Params
id of product
Data Params
None
Headers
Content-Type: application/json

Success Response:
Code: 200
Content:
```
```
{
    "message": "Review has been successfully deleted",
    "success": true
}

``` 

## Backend Development
- REST apis for changing roles of an user and deleting an existing user.
- REST apis for getting,updating and deleting product reviews.
- Bug fix: Negative product stock or product prices will not be accepted
- Bug fix: Check for accepting roles either as a user or an admin has been added
- Bug fix: Check for accepting valid product ratings only (0<=rating<=5) 
- Improved efficiency : Improved average rating calculation by implementing more efficient method

## Test results

### Testing /user/updaterole api

![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/testImages/UpdateUserRole.png)

### Testing /user/delete/{id}

![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/testImages/DeleteUser.png)

### Testing /product/getreviews/{id}

![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/testImages/GetReviews.png)

### Testing /product/updateReview

![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/testImages/UpdateReview.png)

### Testing /product/deleteReview/{id}

![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/testImages/DeleteReview.png)

### Testing bug of accepting invalid role change request

![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/testImages/InvalidRoleChangeRequest.png)

### Testing bug of sending negative product price

![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/testImages/NegativeProductPrice.png)

# Frontend 


## FrontEnd Technologies: React JS, CSS
## Review Product
A layout is created where a user can review a product on the website. The page is mobile responsive. Here are the screenshots to  see where a user can review a product for different screen sizes.
## Order
A layout is created where a user can order the product on the website. The page is mobile responsive. Here are the screenshots to  see where the user can order the product for different screen sizes.

## Cypress Testing
Test Cases are written for different screens.
## Cart
A layout is created where a user can add/delete products in a cart on the website. The page is mobile responsive. Here are the screenshots to  see the checkout page for different screen sizes.



## Test results 
### Cart
![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/testImages/EmptyCart.png)

### Cart-mobile
![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/testImages/EmptyCart_Mobile.png)

### Review
![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/testImages/Review.png)

### Review_Mobile
![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/testImages/Review_Mobile.png)


### Order
![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/testImages/OrderPage.png)
### Order
![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/testImages/OrderPage_Mobile.png)

### Shipping_Details
![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/testImages/Shipping_Details.png)

### CypressTesting
![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/testImages/CypressTesting.png)

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

