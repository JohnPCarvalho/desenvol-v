import React, { Component } from 'react';
import './Travels.css';

import Travel from  './Travel/Travel';
  
  class Travels extends Component {
    constructor(props) {
      super(props);
      this.handleAddition = this.handleAddition.bind(this);
      this.handleSubtraction = this.handleSubtraction.bind(this);

      this.state = {
        users: [
          {
            id: 1,
            driver: "Johnny",
            travelledkilometer: 10.00,
            priceperliter: 4.00,
            literspent: 5,
            checkoutdate: "26/12/1997"
          },
          {
            id: 2,
            driver: "Matheus",
            travelledkilometer: 10.00,
            priceperliter: 4.00,
            literspent: 5,
            checkoutdate: "26/12/1997"
          },
          {
            id: 3,
            driver: "Samuel",
            travelledkilometer: 10.00,
            priceperliter: 4.00,
            literspent: 5,
            checkoutdate: "26/12/1997"
          },
          {
            id: 4,
            driver: "Leonardo",
            travelledkilometer: 10.00,
            priceperliter: 4.00,
            literspent: 5,
            checkoutdate: "26/12/1997"
          },
        ]
      }
    }

    //this wil render the list by mapping the array
    //therefore, the props of Travels component will be an array

    handleAddition (key) {
      alert(key);
    }

    handleSubtraction (key) {
      alert(key);
    }

    render() {

      return (
        <div className="Travels-Container">  
          {
            this.state.users ? 
            this.state.users.map(user => {
              return <Travel
                        id={user.id} 
                        driver={user.driver}  
                        travelledkilometer={user.travelledkilometer}
                        priceperliter={user.priceperliter}
                        literspent={user.literspent}
                        checkoutdate={user.checkoutdate}
                      />
            }) : <div> There is Nothing</div>
          }
        </div>
      )
    }
  
}

export default Travels;