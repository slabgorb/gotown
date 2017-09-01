import React from 'react';
import axios from 'axios';
var Being = require('being.jsx');
var _ = require('underscore');

class Area extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      area: null,
      residents: null,
    }
  }
  render() {
    if (this.state.area) {
      var namesList = _.map(this.state.residents, function(resident, i) {
        return (
          <li key={i}><Being being={resident}/></li>
        )
      })
      return (
        <div className="area" onClick={() => alert('click')}>
          {this.state.area.name}
        <ul className="residents">{namesList}</ul>
        </div>
      )
    } else {
      return <div>Loading...</div>
    }

  }
  componentDidMount() {
    axios.get("/town")
      .then(res => {
        const area = res.data
        this.setState({area:area, residents: area.residents})
      })
  }
}

module.exports = Area
