import React from 'react';
import { Navbar as BSNavbar, Nav } from 'react-bootstrap';

const NavBar = () => {
  return (
    <BSNavbar bg="dark" variant="dark" expand="sm">
      <BSNavbar.Brand href="/">North Star Lodge</BSNavbar.Brand>
      <BSNavbar.Toggle aria-controls="basic-navbar-nav" />
      <BSNavbar.Collapse id="basic-navbar-nav">
        <Nav className="mr-auto">
          <Nav.Link href="/">Home</Nav.Link>
          <Nav.Link href="/maintenance">Maintenance</Nav.Link>
        </Nav>
        <Nav className="ml-auto">
          <Nav.Link href="/sign-out">Sign Out</Nav.Link>
        </Nav>
      </BSNavbar.Collapse>
    </BSNavbar>
  );
};

export default NavBar;
