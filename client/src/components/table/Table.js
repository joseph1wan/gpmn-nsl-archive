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
  {
    firstName: 'Bird',
    lastName: '',
    userName: '@twitter',
  },
];

const Table = (tableData) => {
  tableData = mockTableData;
  const model = tableData[0];

  const keyList = Object.keys(model);

  return (
    <BootstrapTable striped bordered hover>
      <thead>
        <tr>
          <th>#</th>
          {keyList.map((column) => (
            <th>{column}</th>
          ))}
        </tr>
      </thead>
      <tbody>
        {tableData.map((request, index) => (
          <tr>
            <td>{index}</td>
            {keyList.map((column) => (
              <td>{request[column]}</td>
            ))}
          </tr>
        ))}
      </tbody>
    </BootstrapTable>
  );
};

export default Table;
