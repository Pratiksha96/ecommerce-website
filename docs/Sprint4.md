<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Sprint 3](#sprint-3)
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

# Sprint 3

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

Admin:

User:


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
The endpoints developed in this sprint for reviews are - 

```
/product/createReview
/product/getreviews/{id}
/product/updateReview
/product/deleteReview/{id}
```

### PUT /product/createReview

### GET /product/getreviews/{id}

### PUT /product/updateReview

### DELETE /product/deleteReview/{id}

## Backend API Development

## Test results

# Frontend 

## FrontEnd Technologies: React JS, CSS

## Test results 

