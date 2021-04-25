import { checkPropTypes } from 'prop-types';
import React from 'react';
import './TravelsControls.css'

export default function travelsControls( props ) {
    return (
        <div class="Buttons-Container">
            <div>
                <button className="Plus-Button" onClick={props.plus}> + </button>
            </div>
            <div>
                <button className="Minus-Button" onClick={props.minus}> - </button>
            </div>
        </div>
    );
}