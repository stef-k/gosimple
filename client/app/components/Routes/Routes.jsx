import React from 'react';
import  {Router, Route, IndexRoute} from 'react-router';

import {App} from 'App/App';
import {Home} from 'Home/Home';
import {About }from 'About/About';

/**
 * Application's routes
 * @type {Routes}
 */
const Routes = (
  <Router>
    <Route path="/" component={App}>
      <IndexRoute component={Home}/>
      <Route path="/about" component={About}/>
    </Route>
  </Router>
);

export default Routes;
