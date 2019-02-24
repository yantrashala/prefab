import React, { Component } from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import Header from '../header/Header';
import ProjectSetup from './ProjectSetup';
import UISetup from './UISetup';
import MicroservicesSetup from './MicroservicesSetup';
import CiCdSetup from './CiCdSetup';

class GetStarted extends Component {
  render() {
    return (
      <React.Fragment>
        <Header />
        <Route path="/configure/project-setup" component={ProjectSetup} />
        <Route path="/configure/ui-setup" component={UISetup} />
        <Route path="/configure/microservice-setup" component={MicroservicesSetup} />
        <Route path="/configure/cicd-setup" component={CiCdSetup} />
      </React.Fragment>
    );
  }
}

export default GetStarted;
