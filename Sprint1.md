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
- [Adding new products to website](https://github.com/Pratiksha96/ecommerce-website/issues/8)
- [Adding new products to website](https://github.com/Pratiksha96/ecommerce-website/issues/18)


# Backend

## API Contracts 
## Database Schema
## How to run
## Test results

# Frontend 
# FrontEnd Technologies: React JS, CSS

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
![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/sprint1/sprintImages/homepage.png?raw=true)
![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/sprint1/sprintImages/ProductsPage.png?raw=true)
![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/sprint1/sprintImages/MobileResponsive.png?raw=true)
![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/sprint1/sprintImages/Hamburger.png?raw=true)
![alt text](https://github.com/Pratiksha96/ecommerce-website/blob/sprint1/sprintImages/Footer.png?raw=true)

## Note
- On Feb 4, we squashed many commits due to which some of our commits lost their history. User stories closing date will give an idea of our actual commit history.
- Due to git email config issue faced earlier, some commits are shown in two different commit history bar graphs.
