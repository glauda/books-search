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
    <form className={classes.root} noValidate autoComplete="off">
        
        <Grid container spacing={2}>
            <Grid item xs>
                <TextField id="outlined-basic" label="Outlined" variant="outlined" />
            </Grid>
            <Grid item xs>
                <Button color="primary">Search</Button>
            </Grid>
        </Grid>
  
  { /*  <Grid>
            <Row>
                <Cell>
                    <TextField id="outlined-basic" label="Outlined" variant="outlined" />
                </Cell>
                <Cell>
                    <Button color="primary">Search</Button>
                </Cell>
            </Row>
        </Grid>
        */}

    </form>
  );
}