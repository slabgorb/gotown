import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import Typography from '@material-ui/core/Typography';
import PropTypes from 'prop-types';
import React from 'react';
import Expression from '../Chromosome/Expression';
import Chromosome from '../Chromosome/Show';
import beingApi from './api';


const BeingPair = ({ name,id,action }) => {
  return (<div className="being-pair" onClick={() => action(id)}>{name}</div>)
}

BeingPair.propTypes = {
  id: PropTypes.string.isRequired,
  name: PropTypes.string.isRequired,
  action: PropTypes.func.isRequired,
}

const BeingPairs = ({ pairs, action, label }) => {
  return (
    <div>
      <Typography>{label}</Typography>
      { pairs.map(pair => (<BeingPair name={pair.name} id={pair.id} action={action}/>)) }
    </div>
  )
}

BeingPairs.propTypes = {
  pairs: PropTypes.array.isRequired,
  action: PropTypes.func.isRequired,
  label: PropTypes.string,
};

BeingPairs.defaultProps = {
  label: "",
};

class Being extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      being: props.being,
      loaded: false,
    };
    this.handleClick = this.handleClick.bind(this);
  }

  handleClick(value) {
    const { history } = this.props;
    history.push(`/beings/${value}`);
  }

  componentDidMount() {
    this.get(this.props.match.params);
  }

  get({ id }) {
    beingApi.get(id)
      .then((s) => {
        this.setState({
          being: s,
          loaded: true,
        });
      });
  }
  render() {
    const { being, loaded } = this.state;
    if (!loaded) {
      return (<div>Loading</div>);
    }
    const {
      name,
      gender,
      expression,
      chromosome,
      age,
      spouses, 
      children,
      parents,
      species_id, 
      species,
    } = being;
    return (
      <div>
        <Card className="being">
          <CardContent>
            <div className="being-name">{name}</div>
            <div className="being-age">{age}</div>
            <div className="being-gender">{gender}</div>
            <BeingPairs pairs={spouses} label="Spouses" action={this.handleClick} />
            <BeingPairs pairs={children} label="Children" action={this.handleClick} />
            <BeingPairs pairs={parents} label="Parents" action={this.handleClick} />
            <Expression expression={expression} />
            <Chromosome chromosome={chromosome} speciesID={species_id} speciesName={species}/>
          </CardContent>
        </Card>
        <br />
      </div>
    );
  }
}

Being.propTypes = {
  being: PropTypes.object,
  match: PropTypes.object.isRequired,
};

Being.defaultProps = {
  being: {},
};

module.exports = Being;
