import React from 'react';
import { Link } from 'react-router-dom'

class Nav extends React.Component {
  render() {
    return (
      <nav>
        <Link to="Area">Area</Link>
        <Link to="Species">Species</Link>
      </nav>

    )
  }
}

module.exports = Nav
