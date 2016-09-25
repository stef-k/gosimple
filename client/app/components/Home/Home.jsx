import React, {Component} from 'react';
import {Link} from 'react-router';
/**
 * Index page at /
 */
export class Home extends Component {
  render() {
    return (
      <div className="react-container">
        <h3>Hello from react I am a react root (the App) component</h3>
        <p>I am an example of a react component named Home, rendered on / page (App -> Home).</p>
        You can also visit the <Link to="/about">About</Link> page.
      </div>
    )
  }
}
