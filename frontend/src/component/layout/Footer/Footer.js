import React from 'react'
import "./Footer.css"
const Footer = () => {
    return (
        <footer id="footer">
        {/* <div className="leftFooter">
          <h3>About Us</h3>
          <h3>Contact Us</h3>

          <p>Download App for Android and IOS mobile phone</p>
          <img src={playStore} alt="playstore" />
          <img src={appStore} alt="Appstore" />
        </div> */}

        <div className="rightFooter">
          <h4>Contact Us</h4>
          <h4>About Us</h4>

          {/* <a href="*">Instagram</a>
          <a href="*">Facebook</a> */}
        </div>
  
        <div className="midFooter">
          <h1> Team Ehawking</h1>
          <p>High Quality is our first priority</p>
  
          <p>Copyrights 2022 &copy;</p>
        </div>
  
        <div className="rightFooter">
          <h4>Follow Us</h4>
          <a href="*">Instagram</a>
          <a href="*">Facebook</a>
        </div>
      </footer>
    )
}

export default Footer
