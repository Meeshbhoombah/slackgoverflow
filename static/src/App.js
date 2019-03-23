import React, { Component } from 'react';
import { Route } from 'react-router-dom';
import Home from './Components/Home/Home';
import Architect from './Components/Architect/Architect';

class App extends Component {
  render() {
    return (
      <div className="App">
        <Route exact path='/' component={Home} />
        <Route exact path='/architect' component={Architect} />
      </div>
    );
  };
};

export default App;
