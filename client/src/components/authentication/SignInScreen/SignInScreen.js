import React, { useState } from 'react';
import { Container, Form, Button } from 'react-bootstrap';

const SignInScreen = () => {
  const [credentials, setCredentials] = useState({ email: '', password: '' });
  const handleSubmit = (event) => {
    event.preventDefault();
    const {
      elements: {
        signIn: { value: email },
      },
      password: { value: password },
    } = event.target;
    setCredentials({
      email,
      password,
    });
  };
  return (
    <React.Fragment>
      <div className="jumbotron">
        <Container>
          <h1 className="display-4">Welcome to North Star Lodge</h1>
          <hr className="my-4" />
          <p className="lead">Please sign-in with your credentials.</p>
          <code>
            <p>{credentials.email}</p>
            <p>{credentials.password}</p>
          </code>
        </Container>
      </div>
      <Container className="w-50">
        <Form onSubmit={handleSubmit}>
          <Form.Group controlId="signIn">
            <Form.Label>Email Addrress</Form.Label>
            <Form.Control type="email" placeholder="Enter email" />
          </Form.Group>

          <Form.Group controlId="password">
            <Form.Label>Password</Form.Label>
            <Form.Control type="password" placeholder="Password" />
          </Form.Group>

          <Button variant="primary" type="submit" size="lg">
            Submit
          </Button>
        </Form>
      </Container>
    </React.Fragment>
  );
};

export default SignInScreen;
