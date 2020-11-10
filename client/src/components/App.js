import logo from '../logo.svg';
import './App.css';
import React from 'react';
import { Button } from '@material-ui/core';
import ButtonAppBar from './ButtonAppBar.js';
import FormSearch from './Form.js' ;
import DataTable from './ResultTable.js'

import { createMuiTheme, makeStyles, ThemeProvider } from '@material-ui/core/styles';
import teal from '@material-ui/core/colors/teal';
import red from '@material-ui/core/colors/red';

const theme = createMuiTheme({
  palette: {
    primary: {
      main: teal[800],
    },
    secondary: {
      main: red[800],
    },
  },
});


function App() {
  return (
    <ThemeProvider theme={theme}>
      <div className="App" >
        <ButtonAppBar />
        <div style={{"max-width": "750px", margin: 'auto' }}>
          <FormSearch style={{"margin-top": "50px"}} />
          <DataTable />
        </div>
      </div>
    </ThemeProvider>
    
    
  );
}

export default App;
