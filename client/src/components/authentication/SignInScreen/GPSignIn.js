import React from 'react';
import { Container, Button } from 'react-bootstrap';
import { Link } from 'react-router-dom';

const GPSignIn = () => (
  <Container>
    <div className="mb-3">Hey GP Member. GPMail flow is coming soon...</div>
    <Button as={Link} variant="secondary" to="/">
      Return Home
    </Button>
  </Container>
);

export default GPSignIn;
