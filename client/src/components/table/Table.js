import React from 'react';
import { Table as BootstrapTable } from 'react-bootstrap';

const mockTableData = [
  {
    firstName: 'Mark',
    lastName: 'Otto',
    userName: '@mdo',
  },
  {
    firstName: 'Jacob',
    lastName: 'Thornton',
    userName: '@fat',
  },
  {
    firstName: 'Larry the Bird',
    lastName: '',
    userName: '@twitter',
  },
];

const Table = (tableData = mockTableData) => {
  return (
    <BootstrapTable striped bordered hover>
      <thead>
        <tr>
          <th>#</th>
          <th>First Name</th>
          <th>Last Name</th>
          <th>Username</th>
        </tr>
      </thead>
      <tbody>
        {tableData.map((request, index) => (
          <tr>
            <td>{index}</td>
            <td>{request.firstName}</td>
            <td>{request.lastName}</td>
            <td>{request.userName}</td>
          </tr>
        ))}
      </tbody>
    </BootstrapTable>
  );
};

export default Table;
