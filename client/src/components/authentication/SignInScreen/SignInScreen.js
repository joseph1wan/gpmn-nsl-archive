import React, { useState } from 'react';
import { Container, Row, Form, Button } from 'react-bootstrap';

const SignInScreen = () => {
  const [credentials, setCredentials] = useState({ email: '', password: '' });
  const handleSubmit = (event) => {
    event.preventDefault();
    const form = event.target;
    const email = form.elements.signIn.value;
    const password = form.elements.password.value;
    setCredentials({
      email,
      password,
    });
  };
  return (
    <Container>
      <Form onSubmit={handleSubmit}>
        <Row>
          <Form.Group controlId="signIn">
            <Form.Label>Email Addrress</Form.Label>
            <Form.Control type="email" placeholder="Enter email" />
          </Form.Group>
        </Row>

        <Row>
          <Form.Group controlId="password">
            <Form.Label>Password</Form.Label>
            <Form.Control type="password" placeholder="Password" />
          </Form.Group>
        </Row>

        <Row>
          <Button variant="primary" type="submit">
            Submit
          </Button>
        </Row>
      </Form>
      <Row>
        <div>
          Credentials:
          <div>{credentials.email}</div>
          <div>{credentials.password}</div>
        </div>
      </Row>
    </Container>
  );
};

export default SignInScreen;
