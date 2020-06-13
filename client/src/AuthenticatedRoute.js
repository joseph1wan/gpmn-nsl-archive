import React from 'react';
import { Redirect, Route } from 'react-router-dom';
import { useAuthenticationToken } from './hooks';

const AuthenticatedRoute = ({ children, ...rest }) => {
  const [isAuthenticated] = useAuthenticationToken();
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
