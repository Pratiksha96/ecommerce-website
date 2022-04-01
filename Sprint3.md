# Sprint 2

## Stories
Developer:
- [Implementation of User manager should be wrapped under an interface](https://github.com/Pratiksha96/ecommerce-website/issues/99)
Admin:

User:

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

Bug:

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
## API Contracts
The endpoints developed in this sprint for user are - 

```
/me
/password/update
/me/update
/user/get
/user/get/{id}
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

### GET /user/get

### GET /user/get/{id}

### POST /order/create

### GET /order/user/get

### GET /order/get/{id}

### GET /order/get

### DELETE /order/delete/{id}

### PUT /order/update/{id}

## Backend API Development


## Test results


# Frontend 

## Test results 
