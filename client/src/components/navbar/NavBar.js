import React from 'react';
import { Navbar as BSNavbar, Nav } from 'react-bootstrap';
import logo from '../../assets/logo.png';

const NavBar = () => {
  return (
    <BSNavbar bg="dark" variant="dark" expand="sm">
      <BSNavbar.Brand href="/">
        <img
          src={logo}
          width="30"
          height="30"
          className="d-inline-block align-top mr-3"
          alt="North Star Lodge Logo"
        />
        North Star Lodge
      </BSNavbar.Brand>
      <BSNavbar.Toggle aria-controls="basic-navbar-nav" />
      <BSNavbar.Collapse id="basic-navbar-nav">
        <Nav className="ml-auto">
          <Nav.Link href="/info">Info</Nav.Link>
          <Nav.Link href="/maintenance">Maintenance Requests</Nav.Link>
          <Nav.Link href="/exit-checklist">Exit Checklist</Nav.Link>
          <Nav.Link href="/sign-out">Log Out</Nav.Link>
        </Nav>
      </BSNavbar.Collapse>
    </BSNavbar>
  );
};

export default NavBar;
