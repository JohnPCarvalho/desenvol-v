import React, { useState } from 'react';
import './App.css';
import { BrowserRouter, Route, Switch } from 'react-router-dom';

import Dashboard from './components/Dashboard/Dashboard';
import Login from './components/Login/Login';
import Preferences from './components/Preferences/Preferences';
import Travels from './components/Travels/Travels';
import { UserContext } from './contexts/UserContext';

function App() {
  
  const [token, setToken] = useState()

  if(token) {
    return <Login setToken={setToken} />
  }

  return (
    <div className="wrapper">
      <h1>Desenvol-v</h1>
      <BrowserRouter>
        <UserContext.Provider value="hello from context">
          <Switch>
            <Route path="/dashboard">
              <Dashboard />
            </Route>
            <Route path="/preferences">
              <Preferences />
            </Route>
            <Route path="/travels">
              <Travels />
            </Route>
          </Switch>
        </UserContext.Provider>
      </BrowserRouter>
    </div>
  );
}

export default App;
