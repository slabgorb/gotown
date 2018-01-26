import React from 'react';
var _ = require('underscore');
import axios from 'axios';
import Paper from 'material-ui/Paper';
import Typography from 'material-ui/Typography';
import Card, { CardContent } from 'material-ui/Card';
// const Species = (props) => {
//   <SpeciesDisplay { ...props }/>
// }

class Species extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      name: props.name,
      genetics: {traits:[]},
    }
  }
  render() {
    return (
      <div>
        <Paper elevation={4}>
          <Typography type="headline" component="h2">
            {this.state.name}
          </Typography>
          <Genetics traits={this.state.genetics.traits}/>
          <GeneticsMap/>
        </Paper>
      </div>
    )
  }

  componentDidMount() {
    axios.get(`data/${this.props.name}.json`)
      .then(res => {
        const s = res.data;
        this.setState({genetics:s})
      })
  }
}







const GenderName = ({gender, patterns, givenNames}) =>
    (
      <div className="gender-name">
        <h3>{gender}</h3>
        <p>{patterns.join(', ')}</p>
        <p>{givenNames.join(', ')}</p>
      </div>
    )


module.exports = Species
