import React from 'react';
import axios from 'axios';
var Being = require('./Being.jsx');
import Menu, {MenuItem} from '@material-ui/core/Menu';
import List, { ListItem, ListItemText } from '@material-ui/core/List';
var _ = require('underscore');

const sizes = {
  'Hamlet': 5,
  'Village': 7,
  'Town': 8,
  'City': 9,
}

class AreaForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      anchorEl: null,
      open: false,
      selectedIndex: 1,
    }
  }

  handleChange(event) {
    this.setState({ open: true, anchorEl: event.currentTarget });
  }

  render() {

    var optionsList = _.map(sizes, function(i,s) {
      return (<MenuItem key={i} onClick={this.handle}>{s}</MenuItem>)
    });
    return (
      <div>
        <Menu value={this.state.value} onChange={this.handleChange}>
          {optionsList}
        </Menu>
      </div>
    )
  }
}


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
        <div className="area">
          <h1>{this.state.area.name}</h1>
          <ul className="residents">{namesList}</ul>
        </div>
      )
    } else {
      return <div>Loading...</div>
    }

  }
  componentDidMount() {
    axios.get("/town?culture=italian")
      .then(res => {
        const area = res.data
        this.setState({area:area, residents: area.residents})
      })
  }
}

module.exports = {
  Area: Area,
  AreaForm: AreaForm
}
