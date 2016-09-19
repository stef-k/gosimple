'use strict';

var GoSimple = GoSimple || {};

/**
 * Module: GoSimple.main
 * @description
 * @namespace
 * @memberof GoSimple.main
 */
GoSimple.main = (function () {
  // private variables
  var debug = true;
  var websocketSupported = false;
  var serviceUrl = 'localhost:8080';
  var connectTo = 'ws://' + serviceUrl + '/websocket';
  var socket = null;

  // this
  var self = {};
  self.attempts = 1;
  // public variables
  var pub = {};
  pub.connected = false;

  self.debug = function (msg) {
    if (debug) {
      console.debug(msg);
    }
  };

  pub.init = function () {
    self.debug('Simple init');
    self.check();
    self.connect();
  };

  /**
   * Check for websocket support
   */
  self.check = function () {
    if (window.WebSocket) {
      websocketSupported = true
    } else {
      pub.connected = 'GoSimple: Please update to a modern web browser';
    }
  };

  /**
   * Connect to websocket server
   */
  self.connect = function () {
    if (websocketSupported) {
      socket = new WebSocket(connectTo);
      socket.onopen = function () {
        pub.connected = true;
      };
      socket.onmessage = function (e) {
        self.receive(e);
      };
      socket.onclose = function () {
        pub.connected = false;
        self.close();
      };
    }
  };

  /**
   * Handle close event
   */
  self.close = function () {
    self.debug('connection closed, will try to reconnect');
    var time = self.generateInterval(self.attempts);
    setTimeout(function () {
      // increase the attempts by 1
      self.attempts++;

      // try to reconnect
      self.connect();
    }, time);
  };

  // based on http://blog.johnryding.com/post/78544969349/how-to-reconnect-web-sockets-in-a-realtime-web-app
  self.generateInterval = function (k) {
    var maxInterval = (Math.pow(2, k) - 1) * 1000;

    if (maxInterval > 30*1000) {
      maxInterval = 30*1000; // If the generated interval is more than 30 seconds, truncate it down to 30 seconds.
    }

    // generate the interval to a random number between 0 and the maxInterval determined from above
    return Math.random() * maxInterval;
  };

  /**
   * Handle incoming messages
   */
  self.receive = function (e) {
    // print on page
    var el = document.querySelector('#message');
    el.textContent = e.data;
    var msg = JSON.parse(e.data);
    self.debug(msg);
    if (msg.hasOwnProperty('Data')) {
      if (msg.Data.hasOwnProperty('message')) {
        console.log('got message: ', msg.Data.message);
      } else if (msg.Data.hasOwnProperty('cmd')) {
        console.log('got command: ', msg.Data.cmd);
        if (msg.Data.hasOwnProperty('what')) {
          console.log('what: ', msg.Data.what);
        }
      }

    }
  };

  return pub;
})();
document.addEventListener('DOMContentLoaded', GoSimple3.main.init);
