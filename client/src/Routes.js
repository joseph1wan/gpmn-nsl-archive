import React from 'react';
import { Switch, Route } from 'react-router-dom';
import { Container } from 'react-bootstrap';

import App from './App';
import AuthenticatedRoute from './AuthenticatedRoute';
import SignInScreen from './components/authentication/SignInScreen';
import SignOutScreen from './components/authentication/SignOutScreen';
import NavBar from './components/navbar';
import Maintenance from './components/maintenance';

const Routes = () => {
  return (
    <React.Fragment>
      <NavBar />
      <Container>
        <Switch>
          <Route exact path="/">
            <App />
          </Route>

          <AuthenticatedRoute path="/maintenance">
            <Maintenance />
          </AuthenticatedRoute>

          {/* Non-Authenticated Routes */}
          <Route path="/sign-in">
            <SignInScreen />
          </Route>

          <Route path="/sign-out">
            <SignOutScreen />
          </Route>
        </Switch>
      </Container>
    </React.Fragment>
  );
};

export default Routes;
