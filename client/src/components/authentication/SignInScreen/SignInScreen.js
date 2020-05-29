import React, { useState } from 'react';
import { Container } from 'react-bootstrap';
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
      {method === 'gp' ? (
        <GPSignIn />
      ) : (
        <GuestSignIn onSubmit={handleSubmit(onSubmit)} control={control} />
      )}
    </React.Fragment>
  );
};

export default SignInScreen;
