import React, { Component } from 'react';
import './Travels.css';

import Travel from  './Travel/Travel';
  
  class Travels extends Component {
    constructor(props) {
      super(props);
      this.handleAddition = this.handleAddition.bind(this);
      this.handleSubtraction = this.handleSubtraction.bind(this);

      this.state = {
        
      }
      
    }
    //this wil render the list by mapping the array
    //therefore, the props of Travels component will be an array

    handleAddition (key) {
      alert(key)
    }

    handleSubtraction (key) {
      alert(key)
    }

    render() {
      return (
        <div className="Travels-Container">  
          <Travel
            key={1}
            driver="Johnny" 
            travelledkilometer="2.5" 
            priceperliter="4.0" 
            literspent="1" 
            checkoutdate="25/12/1997" 
            minus={() => this.handleSubtraction()}
            plus={() =>this.handleAddition()}
          />
          <Travel
            key={2}
            driver="Johnny" 
            travelledkilometer="2.5" 
            priceperliter="4.0" 
            literspent="1" 
            checkoutdate="25/12/1997" 
            minus={this.handleSubtraction}
            plus={this.handleAddition}
          />
          <Travel
            key={3}
            driver="Johnny" 
            travelledkilometer="2.5" 
            priceperliter="4.0" 
            literspent="1" 
            checkoutdate="25/12/1997" 
            minus={this.handleSubtraction}
            plus={this.handleAddition}
          />
          <Travel
            key={4}
            driver="Johnny" 
            travelledkilometer="2.5" 
            priceperliter="4.0" 
            literspent="1" 
            checkoutdate="25/12/1997" 
            minus={this.handleSubtraction}
            plus={this.handleAddition}
          />
          <Travel
            key={5}
            driver="Johnny" 
            travelledkilometer="2.5" 
            priceperliter="4.0" 
            literspent="1" 
            checkoutdate="25/12/1997" 
            minus={this.handleSubtraction}
            plus={this.handleAddition}
          />
    
        </div>
      )
    }
  
}

export default Travels;