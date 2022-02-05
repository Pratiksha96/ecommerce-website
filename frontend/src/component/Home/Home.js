import React, { Fragment } from 'react'
import Product from './Product.js'
import './Home.css'
import Metadata from '../layout/MetaData'
import axios from 'axios';

const Home = () => {

    async function getData(){
        const data=await axios.get('http://localhost:8081/product/get')
        console.log(data)
    }
    getData();
    
    const product ={
        name:'Blue Tshirt',
        images:[{url:"https://i.ibb.co/DRST11n/1.webp"}],
        price:'$400',
        _id:'aashish'

    }
    return (
        <Fragment>
            <Metadata title="Ecommerce"/>
            <div className="banner">
            <p>Buy for less!</p>
            <h1>Shop Stoppers</h1>

        <a href="#container">
            <button>
                Scroll
                {/* Scroll <CgMouse/> */}
            </button>
        </a>
         </div>
        <h2 className="homeHeading">
            Featured Products
        </h2>
        <div className="container" id="container">
            <Product product={product}/>
            <Product product={product}/>
            <Product product={product}/>
            <Product product={product}/>
            <Product product={product}/>
            <Product product={product}/>
            <Product product={product}/>
            <Product product={product}/>
            <Product product={product}/>
            <Product product={product}/>
            <Product product={product}/>


        </div>

        </Fragment>
    )
}

export default Home
