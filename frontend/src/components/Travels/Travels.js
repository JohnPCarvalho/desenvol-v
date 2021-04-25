import React from 'react';
import './Travels.css';

import Travel from  './Travel/Travel';
  
  const travels = (props) => {
  
    //this wil render the list by mapping the array
    //therefore, the props of Travels component will be an array

    function handleAddition (e) {
      alert(e.target.value);
    }

    function handleSubtraction (e) {
      alert(e.target.value);
    }

  return (
    <div className="Travels-Container">  
      <Travel
        key={1}
        driver="Johnny" 
        travelledkilometer="2.5" 
        priceperliter="4.0" 
        literspent="1" 
        checkoutdate="25/12/1997" 
        minus={handleSubtraction}
        plus={handleAddition}
      />
      <Travel
        key={2}
        driver="Johnny" 
        travelledkilometer="2.5" 
        priceperliter="4.0" 
        literspent="1" 
        checkoutdate="25/12/1997" 
        minus={handleSubtraction}
        plus={handleAddition}
      />
      <Travel
        key={3}
        driver="Johnny" 
        travelledkilometer="2.5" 
        priceperliter="4.0" 
        literspent="1" 
        checkoutdate="25/12/1997" 
        minus={handleSubtraction}
        plus={handleAddition}
      />
      <Travel
        key={4}
        driver="Johnny" 
        travelledkilometer="2.5" 
        priceperliter="4.0" 
        literspent="1" 
        checkoutdate="25/12/1997" 
        minus={handleSubtraction}
        plus={handleAddition}
      />
      <Travel
        key={5}
        driver="Johnny" 
        travelledkilometer="2.5" 
        priceperliter="4.0" 
        literspent="1" 
        checkoutdate="25/12/1997" 
        minus={handleSubtraction}
        plus={handleAddition}
      />

    </div>
  )
}

export default travels;