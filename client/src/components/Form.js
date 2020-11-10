import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import { Button } from '@material-ui/core';
import Grid from '@material-ui/core/Grid';
import Cell from '@material-ui/core/Grid';
import Container from '@material-ui/core/Container';
import Box from '@material-ui/core/Box';


const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
    '& > *': {
      margin: theme.spacing(1),
      width: '25ch',
    },
  },
}));

export default function FormSearch() {
  const classes = useStyles();

  return (
    <div style={{ display: 'inline-flex' , "margin-top": '50px' , "margin-bottom": '20px' }}>
      <div>
        <TextField id="outlined-basic" label="What are you looking for?" variant="outlined" />        
      </div>
      <div style={{ display: 'inline-flex', 'margin-left': '20px'}}>
        <Button color="primary">Search</Button>
      </div>
    </div>
  );
  { /*  
      <form className={classes.root} noValidate autoComplete="off"></form> 
    
    <Grid>
            <Row>
                <Cell>
                    <TextField id="outlined-basic" label="Outlined" variant="outlined" />
                </Cell>
                <Cell>
                    <Button color="primary">Search</Button>
                </Cell>
            </Row>
        </Grid>
        */ }

    
  
}