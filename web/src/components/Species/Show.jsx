import 'typeface-raleway';
import React from 'react';
import axios from 'axios';
import Paper from 'material-ui/Paper';
import Typography from 'material-ui/Typography';
import PropTypes from 'prop-types';
import Genetics from './Genetics';
import GeneticsMap from './GeneticsMap';

class Species extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      name: props.name,
      genetics: { traits: [] },
    };
  }

  componentDidMount() {
    axios.get(`/species/${this.props.name}`)
      .then((res) => {
        const s = res.data;
        this.setState({ genetics: s.expression });
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
          <Genetics traits={this.state.genetics.traits} />
          <GeneticsMap traits={this.state.genetics.traits} />
        </Paper>
      </div>
    );
  }
}

Species.propTypes = {
  name: PropTypes.string.isRequired,
};

module.exports = Species;
