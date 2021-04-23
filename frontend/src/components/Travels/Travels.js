import React from 'react';
import './Travels.css';

import Travel from  './Travel/Travel';
  
  const travels = (props) => {
  
    //this wil render the list by mapping the array
    //therefore, the props of Travels component will be an array

  return (
    <div className="Travels-Container">  
      <div className="Descriptions">
        <p>Driver</p>
        <p>Travelled Km</p>
        <p>Liter spent</p>
        <p>Checkout Date</p>
      </div>
      <Travel driver="Johnny" travelledkilometer="2.5" literspent="1" checkoutdate="25/12/1997" />
      <Travel driver="Johnny" travelledkilometer="2.5" literspent="1" checkoutdate="25/12/1997" />
      <Travel driver="Johnny" travelledkilometer="2.5" literspent="1" checkoutdate="25/12/1997" />
      <Travel driver="Johnny" travelledkilometer="2.5" literspent="1" checkoutdate="25/12/1997" />
    </div>
  )
}

export default travels;