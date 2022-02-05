import React from 'react'
import Header from './component/layout/Header/Header.js'
import Footer from './component/layout/Footer/Footer.js'
import Home from './component/Home/Home.js'
import Contact from './component/Contact/Contact.js';
import About from './component/About/About.js'
import SingleProduct from './component/SingleProduct/SingleProduct'





import {BrowserRouter as Router,Route} from 'react-router-dom'
import WebFont from 'webfontloader';



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
        <Route exact path="/about" component={About}/>
        <Route exact path="/productId" component={SingleProduct}/>


        <Footer/>
    </Router>
  )
}

export default App
