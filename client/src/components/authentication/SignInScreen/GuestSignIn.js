import React from 'react';
import { Container, Form, Button } from 'react-bootstrap';
import { Controller } from 'react-hook-form';
import { Link } from 'react-router-dom';

const GuestSignIn = ({ onSubmit: handleSubmit, control }) => {
  return (
    <Container className="w-50">
      <Form onSubmit={handleSubmit}>
        <Form.Group controlId="signIn">
          <Form.Label>Email Address</Form.Label>
          <Controller
            as={Form.Control}
            type="email"
            placeholder="Enter email"
            name="email"
            control={control}
            defaultValue=""
          />
        </Form.Group>

        <Form.Group controlId="password">
          <Form.Label>Password</Form.Label>
          <Controller
            as={Form.Control}
            type="password"
            placeholder="Password"
            name="password"
            control={control}
            defaultValue=""
          />
        </Form.Group>
        <Container className="text-center">
          <Button className="mr-md-3" variant="primary" type="submit">
            Submit
          </Button>
          <Button className="mt-3 mt-md-0" as={Link} variant="secondary" to="/">
            Return Home
          </Button>
        </Container>
      </Form>
    </Container>
  );
};

export default GuestSignIn;
