import React from 'react';
import { Link } from 'react-router-dom';
import { Container, Button } from 'react-bootstrap';
import logo from './assets/logo.svg';
import './App.scss';

function App() {
  return (
    <div className="App">
      <div className="jumbotron mb-0">
        <Container>
          <h1 className="display-4">Welcome to North Star Lodge</h1>
          <hr className="my-4" />
          <p className="lead">Please select your authorization method below</p>
        </Container>
      </div>
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <Button
          as={Link}
          size="lg"
          to={{ pathname: '/sign-in', method: 'gp' }}
          className="mb-3"
          variant="secondary"
        >
          <code>Gracepoint Member</code>
        </Button>

        <Button as={Link} size="lg" to="/sign-in">
          Guest
        </Button>
      </header>
    </div>
  );
}

export default App;
