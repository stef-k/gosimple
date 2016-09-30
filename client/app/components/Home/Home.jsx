import React, {Component} from 'react';
import {Link} from 'react-router';
import WebsocketComponent from 'Home/Websocket';
/**
 * Index page at /
 */
export class Home extends Component {
  render() {
    return (
      <div className="react-container">
        <h3>WebSocket Viewer (a React container)</h3>
        <p>I am an example of a react component named Home, rendered on / page (App -> Home).</p>
        <p>I am also responsible to render the websocket server status. Every time a new client connects or disconnects, the status
        as seen bellow will change. To check it, open this page on a new tab.</p>
        <WebsocketComponent/>
        You can also visit the <Link to="/about">About</Link> page.
      </div>
    )
  }
}
