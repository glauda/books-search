import logo from '../logo.svg';
import './App.css';
import React from 'react';
import { Button } from '@material-ui/core';
import ButtonAppBar from './ButtonAppBar.js';
import FormSearch from './Form.js' 

function App() {
  return (
    <div className="App">
    {/*
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    */}
      <ButtonAppBar/>
      <FormSearch style={{"margin-top": "50px"}} />
    </div>
    
  );
}

export default App;
