import React from 'react';
import ReactDOM from 'react-dom';
import { Container } from 'react-bootstrap';
import './index.scss';
import * as serviceWorker from './serviceWorker';
import { BrowserRouter } from 'react-router-dom';
import Routes from './Routes';

ReactDOM.render(
  <React.StrictMode>
    <BrowserRouter>
      <Container fluid>
        <Routes />
      </Container>
    </BrowserRouter>
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
