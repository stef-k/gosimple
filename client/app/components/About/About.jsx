import React, {Component} from 'react'
import {Link} from 'react-router';

/**
 * About page at /about
 */
export class About extends Component {
  render() {
    return (
      <div className="react-container">
        <h3>About GoSimple</h3>
      <p>
        About page made exclusively by a React component and routed by react-router.<br/>
        <Link to="/">home</Link>
      </p>
      </div>
    )
  }
}
