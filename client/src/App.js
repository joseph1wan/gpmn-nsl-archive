import React from 'react';
import { Link } from 'react-router-dom';
import logo from './assets/logo.svg';
import './App.scss';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <Link to="/sign-in" className="App-link">
          Sign In Page
        </Link>
      </header>
    </div>
  );
}

export default App;
