import React from 'react';
import { Link } from 'react-router-dom'

class Nav extends React.Component {
  render() {
    return (
      <nav>
        <Link to="area">Area</Link>
        <Link to="species">Species</Link>
      </nav>

    )
  }
}

module.exports = Nav
