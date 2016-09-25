import React, {Component} from 'react';
import {render} from 'react-dom';
import {Router, browserHistory} from 'react-router'

import Routes from 'Routes/Routes';

render(
  <Router  routes={Routes}  history={browserHistory}/>,
  document.getElementById("app"));
