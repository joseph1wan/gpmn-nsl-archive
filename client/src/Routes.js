import React from 'react';
import { Switch, Route } from 'react-router-dom';

import App from './App';
import SignInScreen from './components/authentication/SignInScreen';

const Routes = () => {
  return (
    <div>
      <Switch>
        <Route exact path="/">
          <App />
        </Route>
        <Route path="/sign-in">
          <SignInScreen />
        </Route>
        <Route path="/test" title="Test Page">
          <div>sup</div>
        </Route>
      </Switch>
    </div>
  );
};

export default Routes;
