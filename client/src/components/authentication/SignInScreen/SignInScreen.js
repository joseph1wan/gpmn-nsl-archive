import React, { useState } from 'react';
import { Container, Form, Button } from 'react-bootstrap';
import { useForm, Controller } from 'react-hook-form';
import { Link } from 'react-router-dom';

const SignInScreen = () => {
  const [credentials, setCredentials] = useState({ email: '', password: '' });
  const { handleSubmit, control } = useForm();

  const onSubmit = (data) => {
    setCredentials(data);
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
        <Form onSubmit={handleSubmit(onSubmit)}>
          <Form.Group controlId="signIn">
            <Form.Label>Email Addrress</Form.Label>
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

          <Button variant="primary" type="submit" size="lg">
            Submit
          </Button>

          <Button
            as={Link}
            className="ml-3"
            variant="secondary"
            size="lg"
            to="/"
          >
            Return Home
          </Button>
        </Form>
      </Container>
    </React.Fragment>
  );
};

export default SignInScreen;
