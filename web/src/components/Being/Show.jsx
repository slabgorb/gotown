import React from 'react';
import Card, { CardContent } from 'material-ui/Card';
import PropTypes from 'prop-types';
import Expression from '../Chromosome/Expression';
import Chromosome from '../Chromosome/Show';
import beingApi from './api';


class Being extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      being: props.being,
      loaded: false,
    };
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
    } = being;
    return (
      <div>
        <Card className="being">
          <CardContent>
            <div className="being-name">{name}</div>
            <div className="being-age">{age}</div>
            <div className="being-gender">{gender}</div>
            <Expression expression={expression} />
            <Chromosome chromosome={chromosome} />
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
