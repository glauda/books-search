import logo from '../logo.svg';
import './App.css';
import React from 'react';
import { Button } from '@material-ui/core';
import ButtonAppBar from './ButtonAppBar.js';
import FormSearch from './Form.js' 

function App() {
  return (
    <div className="App">
      <ButtonAppBar/>
      <FormSearch style={{"margin-top": "50px"}} />
    </div>
    
  );
}

export default App;
