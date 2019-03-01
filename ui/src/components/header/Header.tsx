import React from 'react';
import { Link } from 'react-router-dom';
import Grid from '@material-ui/core/Grid';

import '../../App.css';
const Header = () => {
  return (
    <Grid item xs={12}>
      <span className="">&#x25E4;&#x25E3; </span>
      <span className="title"> Prefab | </span>
      <Link to="/configure/project-setup">Project Setup</Link>
      <Link to="ui-setup">UI Setup</Link>
      <Link to="microservice-setup">Microservices Setup</Link>
      <Link to="cicd-setup">CI-CD Setup</Link>
      <hr />
    </Grid>
  );
};
export default Header;
