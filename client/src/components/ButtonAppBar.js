import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';
import IconButton from '@material-ui/core/IconButton';
import MenuIcon from '@material-ui/icons/Menu';
import Link from '@material-ui/core/Link';
import GitHubIcon from '@material-ui/icons/GitHub';
import InvertColorsIcon from '@material-ui/icons/InvertColors';
import { Tooltip } from '@material-ui/core';

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
  },
  menuButton: {
    marginRight: theme.spacing(2),
  },
  title: {
    flexGrow: 1,
  },
}));

export default function ButtonAppBar() {
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <AppBar position="static">
        <Toolbar>
          {/*  
          <IconButton edge="start" className={classes.menuButton} color="inherit" aria-label="menu">
            <MenuIcon />
          </IconButton>
          <Link href="https://github.com/glauda/books-search" color="inherit">Source code</Link>
          */}
          <Tooltip title="Github repository">
            <IconButton href="https://github.com/glauda/books-search" color="inherit" target="_blank" >
              <GitHubIcon  />
            </IconButton>  
          </Tooltip>
          <Typography variant="h4" className={classes.title}>Books search</Typography>
          <Tooltip title="Toogle light/dark theme">
            <IconButton color="inherit" >
              <InvertColorsIcon />
            </IconButton>  
          </Tooltip>
          {/* 
          <Button color="inherit">Login</Button>
          */}
        </Toolbar>
      </AppBar>
    </div>
  );
}