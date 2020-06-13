import React, { useEffect } from 'react';
import { Link, useHistory } from 'react-router-dom';
import { Container, Button, Jumbotron, Row, Col } from 'react-bootstrap';
import { AiOutlineGoogle } from 'react-icons/ai';
import logo from './assets/logo.svg';
import './App.scss';

const App = () => {
  const history = useHistory();
  useEffect(() => {
    const token = localStorage.getItem('authorizationToken');
    // if (!!token) {
    //   history.replace('/maintenance');
    // }
  });
  return (
    <div className="App">
      <Jumbotron className="px-3 mb-0" fluid>
        <Container>
          <h1 className="display-4">Welcome to North Star Lodge</h1>
          <hr className="my-4" />
          <p className="lead">Please select your authorization method below</p>
          <Row>
            <Col>
              <Button
                as={Link}
                size="lg"
                to={{ pathname: '/sign-in', method: 'gp' }}
                className="mr-md-3 mt-md-0 mt-3 text-uppercase align-middle"
                variant="secondary"
              >
                <AiOutlineGoogle
                  className="mr-2 position-relative align-middle"
                  style={{ transform: 'translateY(-2px)' }}
                />
                Sign in with Gpmail
              </Button>

              <Button
                as={Link}
                className="text-uppercase mt-3 mt-md-0"
                size="lg"
                to="/sign-in"
              >
                Guest
              </Button>
            </Col>
          </Row>
        </Container>
      </Jumbotron>
      {/* <div className="App-header">
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
      </div> */}
    </div>
  );
};

export default App;
