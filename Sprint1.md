# Sprint 1 

## Stories
Developer:
- [Server and database connection set up](https://github.com/Pratiksha96/ecommerce-website/issues/3)
- [Server running for application](https://github.com/Pratiksha96/ecommerce-website/issues/4)
- [Backend server low level design](https://github.com/Pratiksha96/ecommerce-website/issues/5)
- [Set up redux for frontend](https://github.com/Pratiksha96/ecommerce-website/issues/9)
- [Header and Footer creation](https://github.com/Pratiksha96/ecommerce-website/issues/14)
- [Get API for single product](https://github.com/Pratiksha96/ecommerce-website/issues/27)
- [Interface implementation for Create API](https://github.com/Pratiksha96/ecommerce-website/issues/29)
- [Interface implementation for Get product API](https://github.com/Pratiksha96/ecommerce-website/issues/32)
- [Interface implementation for Delete and Update APIs](https://github.com/Pratiksha96/ecommerce-website/issues/33)

Admin:
- [Adding new products to website](https://github.com/Pratiksha96/ecommerce-website/issues/7)
- [Deleting invalid products from website](https://github.com/Pratiksha96/ecommerce-website/issues/10)
- [Updating products for website](https://github.com/Pratiksha96/ecommerce-website/issues/11)

User:
- [View products](https://github.com/Pratiksha96/ecommerce-website/issues/8)
- [Landing page for website](https://github.com/Pratiksha96/ecommerce-website/issues/18)


# Backend

## Database Schema
- A database named "ecommerce_website" was created using MongoDB.
- Collection named product was created in this database.
- Schema of this product is as -
```
   { 
       Name        string    
       Description string    
       Price       int       
       Ratings     int       
       Images      []*Image  
       Category    string    
       Stock       int       
       Reviews     []*Review 
   }
```

## API Contracts
We are using http request for operations. The exposed endpoints are - 

```
/ping
/product/get
/product/get/{id}
/product/add
/product/update/{id}
/product/delete/{id}
```

### /ping
Sends a GET request.
Returns a string:

pong

This is only to check if server is running.

### GET /product/get
Returns all users in the system.
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
{
  users: [
           {<user_object>},
           {<user_object>},
           {<user_object>}
         ]
}
``` 

### GET /product/get/{id}
Returns a single product with given pruduct id"
```
URL Params
Product id
Data Params
None
Headers
Content-Type: application/json
Success Response:
Code: 200
Content:
```
Returns product as given - 

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

### POST /product/add
Adds a new product to database. \
Send a POST request with a json body. Eg. 
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
Response - Return the added product.

### PUT /product/update/{id}

Update the product corresponding to product id passed.

```
URL Params
Product id
Data Params
None
Headers
Content-Type: application/json
Success Response:
Code: 200
```
In body Json of object to be updated is passed and update product Json is returned, eg if price is updated to 500-

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

### DELETE /product/delete/{id}

Delete product corresponding to passed product id.

```
URL Params
Product id
Data Params
None
Headers
Content-Type: application/json
Success Response:
Code: 200
```

Response is of type String which tells number of products deleted. eg- 
```
{
    Deleted Count : 1
}
```

## API Development

- REST apis for create, get, update and delete product were made using Go lang. 
- All of the development was done VS code.
- The server was hosted using mux router and the APIs were configured on this.
- APIs were also configured with MongoDB database, and were tested using POSTMAN tool.

## Test results

### Testing product/get api

![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/sprint1/sprintImages/GetAllProducts.png?raw=true)

### Testing product/get/{id} api

![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/sprint1/sprintImages/GetProductByid.png?raw=true)

### Testing product/add api

![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/sprint1/sprintImages/GetAllProducts.png?raw=true)

### Testing product/delete/{id} api

![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/sprint1/sprintImages/DeleteProduct.png?raw=true)

### Testing product/update/{id} api

![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/sprint1/sprintImages/UpdateProduct.png?raw=true)

# Frontend 
## FrontEnd Technologies: React JS, CSS

## Homepage
A layout is created where a header and footer is displayed for all the pages. All the pages are mobile responsive. An hamburger is also created to visit different pages.
Here are the screenshots for the HomePage for different screen sizes.

## Products Page
Products Page is displaying the products with their title, rating, image and price. User gets route to different page on clicking a product. The Product Page is also mobile responsive.

## Connectivity with Backend
- After bringing front end and back end servers up and running, we connected both ends by calling Get all products API from homepage
- Back end data is retrieved successfully by front end but for now we are showing only dummy data 

## Redux Setup
Redux store in setup to store the data fetched from the backend at a single place. So, we can fetch all the data from the store instead of doing prop drilling.

## Test results 

### Home Page
![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/sprint1/sprintImages/homepage.png?raw=true)

### Product Page
![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/sprint1/sprintImages/ProductsPage.png?raw=true)

### Mobile Responsive
![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/sprint1/sprintImages/MobileResponsive.png?raw=true)

### Hamburger
![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/sprint1/sprintImages/Hamburger.png?raw=true)

### Footer
![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/sprint1/sprintImages/Footer.png?raw=true)

### Connectivity between frontend and backend
![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/sprint1/sprintImages/ConnectivityCheck.png?raw=true)

## How to run
### Backend
 - download mongo db - make sure it is running by running command - `mongo`
- checkout main branch and cd into backend
- run - go run main.go start-server 
- It should show in your terminal that Db is connected.
- Verify server is running by running this - localhost:8081/ping - it should result in pong on your screen

### Frontend
 - download node - make sure it is running by running command - `node -v`
- checkout main branch and cd into frontend
- run npm install to install all the dependencies 
- run npm start to start the server
- Verify server is running by running this - localhost:3000 - it should display the home pages 

## Note
- On Feb 4, we squashed many commits due to which some of our commits lost their history. User stories closing date will give an idea of our actual commit history.
- Due to git email config issue faced earlier, some commits are shown in two different commit history bar graphs.

