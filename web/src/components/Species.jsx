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
        console.log(this.state)
        this.setState({genetics:s})
      })
  }
}

const hexes = ['0','1','2','3','4','5','6','7','8','9','a','b','c','d','e','f']
const blankLines = (h) => {
  (
    <div key={h}>
      <span>{h}</span>
      { hexes.map((h) =>  (<span key={h}>&nbsp;</span>))}
    </div>
  )
}

var line = _.map(hexes, function(h) {
  (<span>{h}</span>)
})


const GeneticsMap = () =>
 (
    <div>
      <span>&nbsp;</span>{line}
        { hexes.map((h) => blankLines(h) )}
    </div>
  )

const Trait = (props) =>
  (
    <div>
      <Card>
        <h3>{props.name}</h3>
        {props.variants.map((variant) => (<Variant name={variant.name} match={variant.match} key={variant.name}/>))}
      </Card>
    </div>
  )

const Variant = ({name, match}) =>
  (
      <div className="key-value">
        <div>{name}</div>
        <div>{match}</div>
      </div>
  )

const GenderName = ({gender, patterns, givenNames}) =>
    (
      <div className="gender-name">
        <h3>{gender}</h3>
        <p>{patterns.join(', ')}</p>
        <p>{givenNames.join(', ')}</p>
      </div>
    )

const Genetics = ({traits}) =>
    (
      <div>
        { traits.map((trait) => (<Trait name={trait.name} key={trait.name} variants={trait.variants}/>)) }
      </div>
    )
module.exports = Species
