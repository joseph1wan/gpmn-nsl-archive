import React from 'react';
import { Link } from 'react-router-dom';
import { Container, Button } from 'react-bootstrap';
import { AiOutlineGoogle } from 'react-icons/ai';
import logo from './assets/logo.svg';
import './App.scss';

function App() {
  return (
    <div className="App">
      <div className="jumbotron jumbotron-fluid mb-0">
        <Container>
          <h1 className="display-4">Welcome to North Star Lodge</h1>
          <hr className="my-4" />
          <p className="lead">Please select your authorization method below</p>
        </Container>
      </div>
      <div className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <Button
          as={Link}
          size="lg"
          to={{ pathname: '/sign-in', method: 'gp' }}
          className="mb-3 text-uppercase align-middle"
          variant="secondary"
        >
          <AiOutlineGoogle
            className="mr-2 position-relative align-middle"
            style={{ transform: 'translateY(-2px)' }}
          />
          Sign in with Gpmail
        </Button>

        <Button as={Link} className="text-uppercase" size="lg" to="/sign-in">
          Guest
        </Button>
      </div>
    </div>
  );
}

export default App;
