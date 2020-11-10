import * as React from 'react';
import './ResultTable.css'
import { DataGrid } from '@material-ui/data-grid';

const columns = [
  { field: 'title', headerName: 'Title', width: 200 },
  { field: 'author', headerName: 'Author', width: 200},
  { field: 'headline', headerName: 'Headline', width: 300 },
];

const rows = [
  { id: 1, title: 'Middlemarch', author: 'George Eliot', headline: 'lorem ipsum' },
  { id: 2, title: 'To the Lighthouse', author: 'Virginia Woolf', headline: 'lorem ipsum' },
  { id: 3, title: 'Mrs Dalloway', author: 'Virginia Woolf', headline: 'lorem ipsum' },
  { id: 4, title: 'Great Expectations', author: 'Charles Dickens', headline: 'lorem ipsum' },
  { id: 5, title: 'Jane Eyre', author: 'Charlotte Brontë', headline: 'lorem ipsum' },
  { id: 6, title: 'Bleak House', author: 'Charles Dickens', headline: 'lorem ipsum' },
  { id: 7, title: 'Wuthering Heights', author: 'Emily Brontë', headline: 'lorem ipsum' },
  { id: 8, title: 'David Copperfield', author: 'Charles Dickens', headline: 'lorem ipsum' },
  { id: 9, title: 'Frankenstein', author: 'Mary Shelley', headline: 'lorem ipsum' },
  ];

export default function DataTable() {
  return (
    <div className="DataTable" style={{ height: 400, width: '100%' }}>
      <DataGrid rows={rows} columns={columns} pageSize={5} />
    </div>
  );
}
