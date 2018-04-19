import React from 'react';
import Paper from 'material-ui/Paper';
import Typography from 'material-ui/Typography';
import PropTypes from 'prop-types';
import { ChromosomeShow } from '../Chromosome';
import Genetics from './Genetics';
import GeneticsMap from './GeneticsMap';
import speciesApi from './api';

class Species extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      name: props.match.params.name,
      genetics: { traits: [] },
    };
    this.get = this.get.bind(this);
  }

  componentDidMount() {
    this.get(this.props.match.params);
  }

  componentWillReceiveProps(nextProps) {
    if (this.props.match.params.name !== nextProps.match.params.name) {
      this.get(nextProps.match.params);
    }
  }

  get({ name }) {
    speciesApi.get(name)
      .then((s) => {
        this.setState({ name: s.name, genetics: s.expression });
      });
  }

  render() {
    if (this.state.genetics.traits.length === 0) {
      return (<div />);
    }
    return (
      <div>
        <Paper elevation={4}>
          <Typography variant="headline" component="h1">
            {this.state.name}
          </Typography>
          <ChromosomeShow />
          <Genetics traits={this.state.genetics.traits} />
          <GeneticsMap traits={this.state.genetics.traits} />
        </Paper>
      </div>
    );
  }
}

Species.propTypes = {
  match: PropTypes.object.isRequired,
};

module.exports = Species;
