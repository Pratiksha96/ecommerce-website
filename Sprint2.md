# Sprint 2

## Stories
Developer:
- [Error handling of Creating and Updating product APIs](https://github.com/Pratiksha96/ecommerce-website/issues/46)
- [Error handling of GetAll, Get and Delete product APIs](https://github.com/Pratiksha96/ecommerce-website/issues/47)
- [Products pagination support for better visibility on home page](https://github.com/Pratiksha96/ecommerce-website/issues/52)
- [User model creation](https://github.com/Pratiksha96/ecommerce-website/issues/57)
- [JSON web token support for user login by setting claims](https://github.com/Pratiksha96/ecommerce-website/issues/60)
- [Maintain jwt token within expiration time for user login](https://github.com/Pratiksha96/ecommerce-website/issues/62)
- [Authenticate user using tokens with validations](https://github.com/Pratiksha96/ecommerce-website/issues/63)
- [User token storage for user authentication](https://github.com/Pratiksha96/ecommerce-website/issues/65)
- [Define user and admin as two different roles](https://github.com/Pratiksha96/ecommerce-website/issues/70)

Admin:
- [Admin rights to add, delete and update products](https://github.com/Pratiksha96/ecommerce-website/issues/71)

User:
- [Search products based on product name](https://github.com/Pratiksha96/ecommerce-website/issues/50)
- [Filter products on home page](https://github.com/Pratiksha96/ecommerce-website/issues/51)
- [Register User support](https://github.com/Pratiksha96/ecommerce-website/issues/56)
- [Support for login to e-commerce website](https://github.com/Pratiksha96/ecommerce-website/issues/58)
- [Frontend-Able to register on the website](https://github.com/Pratiksha96/ecommerce-website/issues/82)
- [Frontend-Login and logout on the website](https://github.com/Pratiksha96/ecommerce-website/issues/83)
- [Frontned-Able to search product on the website](https://github.com/Pratiksha96/ecommerce-website/issues/84)
- [Frontend-Able to filter the products based on categories of the product](https://github.com/Pratiksha96/ecommerce-website/issues/85)
- [Frontend-Able to filter the products based on price of the product](https://github.com/Pratiksha96/ecommerce-website/issues/86)
- [Frontend-able to see product details on clicking a product](https://github.com/Pratiksha96/ecommerce-website/issues/87)

Unit Tests:
- [Get All products](https://github.com/Pratiksha96/ecommerce-website/issues/45)
- [Create Product API](https://github.com/Pratiksha96/ecommerce-website/issues/74)
- [Get Product API](https://github.com/Pratiksha96/ecommerce-website/issues/77)

Bug:
- [To be considered as logged-in user as soon as user registers](https://github.com/Pratiksha96/ecommerce-website/issues/79)

# Backend

## Database Schema
- A database named "ecommerce_website" was created using MongoDB.
- Collection named user was created in this database.
- Schema of this user is as -
```
   { 
    Name                string
	Email               string  
	Password            string  
	Avatar              ProfileImage
	Role                string
	ResetPasswordToken  string
	ResetPasswordExpire time.Time
   }
```

## API Contracts
We are using http request for operations. The exposed endpoints are - 

```
/product/search
/register
/login
/logout
```

### GET /product/search?
Returns filtered products in the system based on product name, category, minimum and maximum product prices, page numbers
Sample: http://localhost:8081/product/search?keyword=Green&priceMax=4
```
URL Params
keyword, category, priceMin, priceMax, page
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
        "_id": "62227cefdf258bf931110624",
        "category": "Category8",
        "description": "Description1",
        "images": [
            {
                "public_id": "id8",
                "url": "url8"
            }
        ],
        "name": "Product1",
        "price": 100,
        "ratings": 5,
        "reviews": null,
        "stock": 10
    },
    {
        "_id": "62227cf8df258bf931110625",
        "category": "Category8",
        "description": "Description2",
        "images": [
            {
                "public_id": "id8",
                "url": "url8"
            }
        ],
        "name": "Product2",
        "price": 700,
        "ratings": 5,
        "reviews": null,
        "stock": 10
    },
    {
        "_id": "62227d03df258bf931110626",
        "category": "Category8",
        "description": "Description3",
        "images": [
            {
                "public_id": "id8",
                "url": "url8"
            }
        ],
        "name": "Product3",
        "price": 5200,
        "ratings": 5,
        "reviews": null,
        "stock": 10
    },
    {
        "_id": "62227d0adf258bf931110627",
        "category": "Category8",
        "description": "Description4",
        "images": [
            {
                "public_id": "id8",
                "url": "url8"
            }
        ],
        "name": "Product4",
        "price": 7500,
        "ratings": 5,
        "reviews": null,
        "stock": 10
    },
    {
        "_id": "62227d11df258bf931110628",
        "category": "Category8",
        "description": "Description5",
        "images": [
            {
                "public_id": "id8",
                "url": "url8"
            }
        ],
        "name": "Product5",
        "price": 10000,
        "ratings": 5,
        "reviews": null,
        "stock": 10
    },
    {
        "_id": "62227d19df258bf931110629",
        "category": "Category8",
        "description": "Description6",
        "images": [
            {
                "public_id": "id8",
                "url": "url8"
            }
        ],
        "name": "Product6",
        "price": 200,
        "ratings": 5,
        "reviews": null,
        "stock": 10
    },
    {
        "_id": "62227d24df258bf93111062a",
        "category": "Category8",
        "description": "Description7",
        "images": [
            {
                "public_id": "id8",
                "url": "url8"
            }
        ],
        "name": "Product7",
        "price": 201,
        "ratings": 5,
        "reviews": null,
        "stock": 10
    },
    {
        "_id": "62227d2bdf258bf93111062b",
        "category": "Category8",
        "description": "Description8",
        "images": [
            {
                "public_id": "id8",
                "url": "url8"
            }
        ],
        "name": "Product8",
        "price": 1900,
        "ratings": 5,
        "reviews": null,
        "stock": 10
    },
    {
        "_id": "62227d32df258bf93111062c",
        "category": "Category8",
        "description": "Description9",
        "images": [
            {
                "public_id": "id8",
                "url": "url8"
            }
        ],
        "name": "Product9",
        "price": 7100,
        "ratings": 5,
        "reviews": null,
        "stock": 10
    },
    {
        "_id": "62227d39df258bf93111062d",
        "category": "Category8",
        "description": "Description10",
        "images": [
            {
                "public_id": "id8",
                "url": "url8"
            }
        ],
        "name": "Product10",
        "price": 9200,
        "ratings": 5,
        "reviews": null,
        "stock": 10
    }
]
``` 

### POST /register
Returns User token and status.
```
URL Params
None
Data Params
{
    "name" : "Test2",
	"email" : "test2@gmail.com",
	"password"  : "testUser1",
    "avatar" :{
        "public_id" : "id2",
        "url" : "url2"
    }
}
Headers
Content-Type: application/json

Success Response:
Code: 200
Content:
```
```
{
    "success": true,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InN3YXRpMjJAZ21haWwuY29tIiwiZXhwIjoxNjQ2NDI3ODY2fQ.x1x1HsLnYmme8XbmTjx0ozscGFwSDhqj49hDrPMdlpo"
}
``` 

### POST /login
Returns User token and status.
```
URL Params
None
Data Params
{
	"email" : "test2@gmail.com",
    "password"  : "testUser1"
}
Headers
Content-Type: application/json

Success Response:
Code: 200
Content:
```
```
{
    "success": true,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InN3YXRpMjJAZ21haWwuY29tIiwiZXhwIjoxNjQ2NDI3OTE0fQ.gouIqZu47jgWhsB8Nti-ZBkeTmJXfOdvA9sccARet5M"
}
``` 

### POST /logout
Removes User token and cleans cookie.
```
URL Params
None
Data Params
{
	"email" : "test2@gmail.com",
    "password"  : "testUser1"
}
Headers
Content-Type: application/json

Success Response:
Code: 200
Content:
```
```
Removes token from the cookies for the respective user
``` 

## Backend API Development

- REST apis for search and filter products with pagination, user register, login and logout using Go lang.
- Search and filteration is done on the basis of product names, categories, price minimum and max value, page numbers
- All of the development was done VS code.
- The server was hosted using mux router and the APIs were configured on this.
- APIs were also configured with MongoDB database using product and user collections, and were tested using POSTMAN tool.

## Test results
### Testing product/search/keyword={name} api

![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/SearchProductByName.png)

### Testing product/search/keyword={name}&category={category} api

![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/SearchProductByNameAndCategory.png)

### Testing product/search/keyword={name}&priceMin={priceMin} api

![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/SearchProductByNameAndPriceMin.png)

### Testing product/search/keyword={name}&priceMax={priceMax} api

![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/SearchProductByNameAndPriceMax.png)

### Testing product/search/page={page} api

![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/Pagination.png)


### Testing /register api

![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/RegisterUser.png)

### Testing /login api

![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/LoginUser.png)


### Testing /logout api

![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/LogoutUser.png)


# Frontend 
## FrontEnd Technologies: React JS, CSS

## Register User
A layout is created where a user can register itself on the website. The page is mobile responsive. Here are the screenshots to  Register the user for different screen sizes.

## Login User
A layout is created where a user can login and logout from the website. The page is mobile responsive. Here are the screenshots to  login the user for different screen sizes.

## Search Products
- A layout is created where a user can search a product from the website. 
- All the related products will be displayed on the frontend.

## Filter Products on Category
- A layout is created where a user can filter the products based on the category  from the website. 
- All the related products will be displayed on the frontend.

## Filter Products on Price
- A layout is created where a user can filter the products based on the price  from the website. 
- All the related products will be displayed on the frontend.

## Single Product Page
- A layout is created where a single product details and description are displayed on the frontend.

## Cypress testing
- To perform the integration testing, test cases are written in cypress.


## Test results 
### HomePage Register User
![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/HomepageRegister.png)

### HomePage Login User
![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/HomepageLogin.png)

### HomePage Search by name
![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/HomepageSearchByName.png)

### HomePage Search by category
![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/HomepageSearchByCategory.png)

### Front end Cypress
![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/main/resources/FrontendCypress.png)


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
q