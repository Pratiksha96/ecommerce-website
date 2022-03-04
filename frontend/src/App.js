import React from 'react'
import Header from './component/layout/Header/Header.js'
import Footer from './component/layout/Footer/Footer.js'
import Loader from './component/layout/Loader/Loader.js'

import Home from './component/Home/Home.js'
import Contact from './component/Contact/Contact.js';
import About from './component/About/About.js'
import SingleProduct from './component/SingleProduct/SingleProduct'
import ProductDetails from './component/Product/ProductDetails.js'
import {BrowserRouter as Router,Route} from 'react-router-dom'
import WebFont from 'webfontloader';
import Products from './component/Product/Products.js'
import Search from './component/Product/Search.js'
import LoginSignUp from './component/User/LoginSignUp.js'


const App = () => {
  React.useEffect(()=>{
    WebFont.load({
      google:{
        families:["Papyrus", "Droid Sans","Chilanka"]
      }
    })
    },[]);
  return (
    <Router>
        <Header/> 
        <Route exact path="/" component={Home}/>
        <Route exact path="/contact" component={Contact}/>
        <Route exact path="/products" component={Products}/>
        <Route  path="/products/:keyword" component={Products}/>
        <Route  path="/login" component={LoginSignUp}/>


        <Route exact path="/search" component={Search}/>


        <Route exact path="/about" component={About}/>
        <Route exact path="/productId" component={SingleProduct}/>
        <Route exact path="/product/:id" component={ProductDetails}/>


        

        <Footer/>
    </Router>
  )
}

export default App
