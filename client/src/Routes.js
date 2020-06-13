import React, { useEffect } from 'react';
import { Switch, Route, useHistory } from 'react-router-dom';

import App from './App';
import AuthenticatedRoute from './AuthenticatedRoute';
import SignInScreen from './components/authentication/SignInScreen';
import SignOutScreen from './components/authentication/SignOutScreen';
import NavBar from './components/navbar';
import Maintenance from './components/maintenance';

const Routes = () => {
  return (
    <div>
      <NavBar />
      <Switch>
        <Route exact path="/">
          <App />
        </Route>
        <Route path="/sign-in">
          <SignInScreen />
        </Route>
        <AuthenticatedRoute path="/maintenance">
          <Maintenance />
        </AuthenticatedRoute>
        <Route path="/sign-out">
          <SignOutScreen />
        </Route>
      </Switch>
    </div>
  );
};

export default Routes;
