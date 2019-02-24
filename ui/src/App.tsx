import React, { Component } from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import './App.css';
import Home from './components/Home';
import GetStarted from './components/configure/GetStarted';
class App extends Component {
  render() {
    return (
      <div className="container">
        <BrowserRouter>
          <Switch>
            <Route path="/" component={Home} exact />
            <Route path="/configure" component={GetStarted} />
          </Switch>
        </BrowserRouter>
      </div>
    );
  }
}

export default App;
