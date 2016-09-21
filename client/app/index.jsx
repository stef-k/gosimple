import React, { Component } from 'react';
import ReactDOM from 'react-dom';

import {ExampleComponent} from 'Example/Example';

/* Application's root element */
class App extends Component {
  render() {
    return (
      <div>
        <h3>Hello from react</h3>
        <ExampleComponent/>
      </div>
    )
  };
}

ReactDOM.render(
  <App/>,
  document.getElementById('app')
);
