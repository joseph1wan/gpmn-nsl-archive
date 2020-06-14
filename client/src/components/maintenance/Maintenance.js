import React from 'react';
import { Jumbotron, Container } from 'react-bootstrap';
import Table from '../table';

import './maintenance.scss';

const Maintenance = () => {
  return (
    <React.Fragment>
      <div className="banner text-center">
        <h1 className="banner--title text-light font-weight-light">
          North Star Lodge
        </h1>
      </div>
      <Jumbotron fluid>
        <Container>
          <h1 className="display-4">Maintenance Requests</h1>
          <hr className="my-4" />
          <p className="lead">Review or submit new Maintenance Requests</p>
        </Container>
      </Jumbotron>
      <Table />
    </React.Fragment>
  );
};

export default Maintenance;
