import React from 'react';
import './Travels.css';

import Travel from  './Travel/Travel';
  

  const travels = (props) => {
  
  return (
    <div>
      <ul>
        <li><Travel driver="Johnny" travelledkilometer="2.5" literspent="1" checkoutdate="25/12/1997" /> </li>
      </ul>
    </div>
  )
}

export default travels;