import React, { useState } from 'react';
import './Login.css'

export default function Login() {


  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  return (
    <div className="login-wrapper">
      <h1>Please log in</h1>
      <form>
        <label>
          <p>Username</p>
          <input type="text"/>
        </label>
        <label>
          <p>Password</p>
          <input type="password"/>
        </label>
        <div>
          <button type="submit">Submit</button>
        </div>
      </form>
    </div>
  )
}