import React, { useState } from 'react';
import Button from '@material-ui/core/Button';
import { MuiThemeProvider, withStyles } from '@material-ui/core/styles';
import psTheme from '../../themes/ps-theme';
import MenuItem from '@material-ui/core/MenuItem';
import TextField from '@material-ui/core/TextField';
import '../../App.css';

const frameworks = [
  {
    value: 'csr',
    label: 'Client Rendered SPA'
  },
  {
    value: 'ssr',
    label: 'Server Rendered Isomorphic '
  },
  { value: 'microfrontend', label: 'Microfrontend ' }
];

const jsLanguages = [
  {
    value: 'es6',
    label: 'JavaScript-ES6+'
  },
  {
    value: 'ts',
    label: 'TypeScript'
  }
];

const cssStyles = [
  {
    value: 'scss',
    label: 'SAS/SCSS'
  },
  {
    value: 'styled-components',
    label: 'Styled Components'
  },
  {
    value: 'emotion-js',
    label: 'Emotion JS'
  }
];

const UISetup = (props: any) => {
  const [architecture, setArchitecture] = useState('csr');
  const [jsLanguage, setJsLanguage] = useState('js');
  const [cssStyle, setCssStyle] = useState('scss');
  const { classes } = props;

  return (
    <MuiThemeProvider theme={psTheme}>
      <h1>UI Setup</h1>
      <form className="form-container">
        <TextField
          id="architecture"
          select
          fullWidth
          label="Select your UI Architecture"
          value={architecture}
          onChange={() => setArchitecture(architecture)}
          margin="normal"
        >
          {frameworks.map(option => (
            <MenuItem key={option.value} value={option.value}>
              {option.label}
            </MenuItem>
          ))}
        </TextField>
        <TextField
          id="jsLanguage"
          select
          fullWidth
          label="Select JavaScript Language"
          value={jsLanguage}
          onChange={() => setJsLanguage(jsLanguage)}
          margin="normal"
        >
          {jsLanguages.map(option => (
            <MenuItem key={option.value} value={option.value}>
              {option.label}
            </MenuItem>
          ))}
        </TextField>
        <TextField
          id="css-styles"
          select
          fullWidth
          label="Select Styling Options"
          value={cssStyle}
          onChange={() => setCssStyle(cssStyle)}
          helperText="Please select Styling Options"
          margin="normal"
        >
          {jsLanguages.map(option => (
            <MenuItem key={option.value} value={option.value}>
              {option.label}
            </MenuItem>
          ))}
        </TextField>

        <Button variant="contained" color="primary">
          Setup UI & Continue{' '}
        </Button>
      </form>
    </MuiThemeProvider>
  );
};

export default UISetup;
