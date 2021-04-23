import React from 'react';
import './Travel.css';

const travel = (props) => {
  return (
    <div>
      <p>{props.driver}</p>
      <p>{props.travelledkilometer} </p>
      <p>{props.literspent} </p>
      <p>{props.checkoutdate} </p>
    </div>
  );
}

export default travel;