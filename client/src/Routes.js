import React from 'react';
import { BrowserRouter, Switch, Route, Link } from 'react-router-dom';

import App from './App';

const Routes = () => {
  return (
    <div>
      <Route exact path="/">
        <App />
      </Route>
    </div>
  );
};

export default Routes;
