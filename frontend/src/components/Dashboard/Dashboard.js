import React, { useContext } from 'react';
import { Button, Container } from 'reactstrap';
import { Link } from 'react-router-dom';
import { UserContext } from '../../contexts/UserContext';

export default function Dashboard() {

  const msg = useContext(UserContext);

  return (  
    <div>
      <h2>Dashboard</h2>
      <div>{msg}</div>
      <Link to="/travels"><Button>Travels</Button></Link>
      <Link to="/preferences"> <Button>Preferences</Button></Link>
    </div>
  )
}