import React, { useContext, useState } from 'react';
import PropTypes from 'prop-types'
import './Login.css'
import api from '../../api/api';
import { Context } from '../../Context/AuthContext';


export default function Login({ setToken }) {

  const authenticated = useContext(Context);
  console.debug('Login', authenticated)

  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  function handleSubmit(event) {
    api.post('/login', 
      {
        email: email,
        password: password
      },
    {withCredentials: false}
    )
    .then(response => {
      console.log("res from login", response);
    }).catch(error => {
      console.log("login error: ", error)
    });
    event.preventDefault();
    
  }

  return (
    <div className="login-wrapper">
      <h1>Please log in</h1>
      <form onSubmit={handleSubmit}>
        <label>
          <p>Username</p>
          <input name="email" type="text" value={email} onChange={e => setEmail(e.target.value)}/>
        </label>
        <label>
          <p>Password</p>
          <input name="password" type="password" value={password} onChange={e => setPassword(e.target.value)}/>
        </label>
        <div>
          <button type="submit">Submit</button>
        </div>
      </form>
    </div>
  )
}

Login.propTypes = {
  setToken: PropTypes.func.isRequired
}
