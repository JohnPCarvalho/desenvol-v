import React from 'react';
import { Switch, Route } from 'react-router-dom';

import Login from './components/Login/Login';
//importar dashboard
//importar travels
import Preferences from './components/Preferences/Preferences';

export default function Routes() {
    return (
        <Switch>
            <Route exact path="/login" component={Login}/>
            <Route exact path="/preferences" component={Preferences}/>
        </Switch>
    );
}