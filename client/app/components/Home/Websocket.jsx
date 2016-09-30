import React, {Component} from 'react';
import incommingMessage from 'Actions/IncomingMessage'
import store from 'Reducers/index'
/**
 * Websocket class to be used by a React component
 * which will be dispatch Redux actions to the store
 */
class Websocket {
  constructor(connectionString, store) {
    this.connectionString = connectionString;
    this.connected = false;
    this.socket = null;
    this.browserOk = false;
    this.checkBrowser();
  }

  /**
   * Checks browser compatibility for the WebSocket protocol
   */
  checkBrowser() {
    if (window.WebSocket) {
      this.browserOk = true
    } else {
      this.browserOk = false;
      console.log('GoSimple: Please update to a modern web browser');
    }
  }

  /**
   * Creates the connection and sets the event handlers
   */
  connect() {
    if (this.browserOk) {
      this.socket = new WebSocket(this.connectionString);

      this.socket.onopen = () => {
        this.connected = true;
      };
      this.socket.onmessage = (e) => {
        this.receive(e);
      };

      this.socket.onclose = () => {
        this.connected = false;
        this.reconnect();
      };
    }
  }

  /**
   * Reconnects the websocket connection in case of communication failure
   */
  reconnect() {
    let time = this.generateInterval(this.attempts);
    setTimeout(function () {
      // increase the attempts by 1
      this.attempts++;
      // try to reconnect
      this.connect();
    }.bind(this), time);

  };

  // based on http://blog.johnryding.com/post/78544969349/how-to-reconnect-web-sockets-in-a-realtime-web-app
  generateInterval(k) {
    let maxInterval = (Math.pow(2, k) - 1) * 1000;

    if (maxInterval > 30 * 1000) {
      maxInterval = 30 * 1000; // If the generated interval is more than 30 seconds, truncate it down to 30 seconds.
    }

    // generate the interval to a random number between 0 and the maxInterval determined from above
    return Math.random() * maxInterval;
  };

  /**
   * Handle incoming messages
   */
  receive(e) {
    console.log('Got message from WebSocket: ', e.data);
    store.dispatch(incommingMessage(e.data))
  };
}

/**
 * A React component responsible to set up a WebSocket connection
 * and link it to Redux store dispatcher.
 * @see 'Actions/IncomingMessages' for the action creator
 */
export default class WebsocketComponent extends Component {

  constructor() {
    super();
    // get address and port from the injected index.html file
    this.address = GoSimple.vars.address;
    this.port = GoSimple.vars.port;
    this.websocket = new Websocket('ws://' + this.address + ':' + this.port + '/websocket');
  }

  componentDidMount() {
    store.subscribe(() => this.forceUpdate());
    this.websocket.connect();
  }

  render() {
    const message = store.getState().message;
    return (
      <div>
        <h4>Incoming message:</h4>
        <pre>
        {message.text}
      </pre>
      </div>
    );
  }
}
