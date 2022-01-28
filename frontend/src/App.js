import React from 'react'
import Header from './component/layout/Header/Header.js'
import Footer from './component/layout/Footer/Footer.js'


import {BrowserRouter as Router} from 'react-router-dom'
import WebFont from 'webfontloader';



const App = () => {
  React.useEffect(()=>{
    WebFont.load({
      google:{
        families:["Roboto", "Droid Sans","Chilanka"]
      }
    })
    },[]);
  return (
    <Router>
        <Header/> 
        <Footer/>
    </Router>
  )
}

export default App
