import React from 'react';
import { Button, Container } from 'reactstrap';
import { Link } from 'react-router-dom';

export default function Dashboard() {
  return (  
    <div>
      <h2>Dashboard</h2>
      <Link to="/travels"><Button>Travels</Button></Link>
      <Link to="/preferences"> <Button>Preferences</Button></Link>
    </div>
  )
}