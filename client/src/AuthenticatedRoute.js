import React from 'react';
import { Redirect, Route } from 'react-router-dom';

const AuthenticatedRoute = ({ children, ...rest }) => {
  const token = localStorage.getItem('authorizationToken');
  const isAuthenticated = !!token;
  return (
    <Route
      {...rest}
      render={() => {
        if (!isAuthenticated) return <Redirect to="/" />;
        else return children;
      }}
    />
  );
};

export default AuthenticatedRoute;
