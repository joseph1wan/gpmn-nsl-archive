import React, { useState } from 'react';
import { Container, Jumbotron } from 'react-bootstrap';
import { useForm } from 'react-hook-form';
import { useLocation } from 'react-router-dom';

import GPSignIn from './GPSignIn';
import GuestSignIn from './GuestSignIn';

const SignInScreen = () => {
  const [credentials, setCredentials] = useState({ email: '', password: '' });
  const { handleSubmit, control } = useForm();
  const { method } = useLocation();

  const onSubmit = (data) => {
    setCredentials(data);
  };
  return (
    <React.Fragment>
      <Jumbotron fluid className="px-3">
        <Container>
          <h1 className="display-4">Welcome to North Star Lodge</h1>
          <hr className="my-4" />
          <p className="lead">Please sign-in with your credentials.</p>
          <code>
            <p>{credentials.email}</p>
            <p>{credentials.password}</p>
          </code>
        </Container>
      </Jumbotron>
      {method === 'gp' ? (
        <GPSignIn />
      ) : (
        <GuestSignIn onSubmit={handleSubmit(onSubmit)} control={control} />
      )}
    </React.Fragment>
  );
};

export default SignInScreen;
