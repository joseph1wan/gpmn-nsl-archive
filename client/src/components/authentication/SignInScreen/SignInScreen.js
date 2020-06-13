import React, { useState } from 'react';
import axios from 'axios';
import { Container, Jumbotron } from 'react-bootstrap';
import { useForm } from 'react-hook-form';
import { useLocation, useHistory } from 'react-router-dom';

import GPSignIn from './GPSignIn';
import GuestSignIn from './GuestSignIn';

const SignInScreen = () => {
  const [credentials, setCredentials] = useState({ email: '', password: '' });
  const { handleSubmit, control } = useForm();
  const { method } = useLocation();
  const history = useHistory();

  const onSubmit = (data) => {
    setCredentials(data);
    axios
      .post('/login', credentials)
      .then((response) => {
        const { token } = response;
        localStorage.setItem('authorizationToken', token);
        history.replace('/');
      })
      .catch((error) => {
        console.log(error);
      });
    localStorage.setItem('authorizationToken', 'authorized');
    history.replace('/');
  };
  return (
    <React.Fragment>
      <Jumbotron fluid className="px-3">
        <Container>
          <h1 className="display-4">Welcome to North Star Lodge</h1>
          <hr className="my-4" />
          <p className="lead">Please sign-in with your credentials.</p>
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
