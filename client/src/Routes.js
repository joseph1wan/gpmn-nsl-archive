import React, { useEffect } from 'react';
import { Switch, Route, useHistory } from 'react-router-dom';

import App from './App';
import SignInScreen from './components/authentication/SignInScreen';
import SignOutScreen from './components/authentication/SignOutScreen';
import NavBar from './components/navbar';
import Maintenance from './components/maintenance';

const Routes = () => {
  const history = useHistory();
  useEffect(() => {
    const token = localStorage.getItem('authorizationToken');
    if (token === null) {
      history.replace('/');
    }
  });
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
        <Route path="/maintenance">
          <Maintenance />
        </Route>
        <Route path="/sign-out">
          <SignOutScreen />
        </Route>
      </Switch>
    </div>
  );
};

export default Routes;
