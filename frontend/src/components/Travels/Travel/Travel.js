import React from 'react';
import './Travel.css';

import Aux from '../../../hoc/Aux';
import TravelsControls from '../TravelsControls/TravelsControls'

const travel = (props) => {
  return (
    <Aux>
      <div className="Travel-Container">
        <p>{props.driver}</p>
        <p>{props.travelledkilometer} </p>
        <p>R$ {props.priceperliter}</p>
        <p>{props.literspent} </p>
        <p>{props.checkoutdate}</p>
        <TravelsControls plus={props.plus} plus={props.minus} />
      </div>

    </Aux>

  );
}

export default travel;