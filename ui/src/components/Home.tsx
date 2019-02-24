import React from 'react';
import Button from '@material-ui/core/Button';
import { MuiThemeProvider } from '@material-ui/core/styles';
import psTheme from '../themes/ps-theme';
import { Link } from 'react-router-dom';

const Home = () => {
  return (
    <MuiThemeProvider theme={psTheme}>
      <header className="App-header">
        <h1>&#x25E4;&#x25E3;</h1>
        <h2>prefab</h2>
        <small>fastest way to get started</small>
        <Button variant="outlined" color="primary">
          Get started
        </Button>
        <Link to="/configure">Get Started</Link>
      </header>
    </MuiThemeProvider>
  );
};
export default Home;
