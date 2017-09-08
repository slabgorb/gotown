import React from 'react';
var _ = require('underscore');
import axios from 'axios';
import Paper from 'material-ui/Paper';
import Typography from 'material-ui/Typography';
import Card, { CardContent } from 'material-ui/Card';

class Species extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      name: "Loading...",
      genetics: {},
      matronymics: [],
      patronymics: [],
      genderNames: [],
    }
  }
  render() {
    var genderNames =_.map(this.state.genderNames, function(gn) {
      return (<GenderName key={gn.gender} gender={gn.gender} patterns={gn.patterns} givenNames={gn.given_names}/>)
    })
    return (
      <div>
        <Paper elevation={4}>
          <Typography type="headline" component="h2">
            {this.state.name}
          </Typography>
          <Genetics traits={this.state.genetics.traits}/>
          {genderNames}
        </Paper>
      </div>
    )
  }

  componentDidMount() {
    axios.get(`data/${this.props.name}.json`)
      .then(res => {
        const s = res.data;
        this.setState({name: s.name, genetics: s.genetics, genderNames: s.gender_names})
      })
  }
}

class Trait extends React.Component {
  render() {
    var variants = _.map(this.props.variants, function(variant) {
      return(<Variant name={variant.name} match={variant.match} key={variant.name}/>)
    });

    return (
      <div>
        <Card>
          <h3>{this.props.name}</h3>
          {variants}
        </Card>
      </div>
    )
  }
}

class Variant extends React.Component {
  render() {
    return (
        <div className="key-value">
          <div>{this.props.name}</div>
          <div>{this.props.match}</div>
        </div>
    )
  }
}

class GenderName extends React.Component {
  render() {
    return (
      <div className="gender-name">
        <h3>{this.props.gender}</h3>
        <p>{this.props.patterns.join(', ')}</p>
        <p>{this.props.givenNames.join(', ')}</p>
      </div>
    )
  }
}

class Genetics extends React.Component {
  render() {
    return (
      <div>
        {
          _.map(this.props.traits, function(trait) {
            return(<Trait name={trait.name} key={trait.name} variants={trait.variants}/>)
          })
        }
      </div>
    )
  }
}

module.exports = Species
