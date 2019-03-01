import React from 'react';
import Button from '@material-ui/core/Button';
import { MuiThemeProvider, withStyles } from '@material-ui/core/styles';
import psTheme from '../../themes/ps-theme';
import { Link } from 'react-router-dom';
import TextField from '@material-ui/core/TextField';
import '../../App.css';
import Grid from '@material-ui/core/Grid';
//-- imports for the editor
// import Editor from 'react-simple-code-editor';
// import { highlight, languages } from 'prismjs/components/prism-core';
// import 'prismjs/components/prism-clike';
// import 'prismjs/components/prism-javascript';

//-- imports for the editor

const ProjectSetup = (props: any) => {
  const { classes } = props;
  return (
    <MuiThemeProvider theme={psTheme}>
      <Grid item xs={6}>
        <h1>Project Setup</h1>
        <form className="form-container">
          <TextField fullWidth color="secondary" id="project-id" label="Project ID" margin="normal" className="textField" variant="filled" />
          <TextField fullWidth id="project-name" label="Project Name" margin="normal" className="textField" variant="filled" />
          <TextField fullWidth id="engineering-lead" label="Engineering Lead Name" margin="normal" className="textField" variant="filled" />
          <TextField fullWidth id="git-url" label="Git URL" margin="normal" className="textField" variant="filled" />
          <TextField fullWidth id="jira-url" label="Jira URL" margin="normal" className="textField" variant="filled" />
          <Button variant="contained" color="primary">
            Setup Project & Continue{' '}
          </Button>
        </form>
      </Grid>
    </MuiThemeProvider>
  );
};
export default ProjectSetup;
