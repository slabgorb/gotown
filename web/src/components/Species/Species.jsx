import React from 'react';
import axios from 'axios';
import Paper from 'material-ui/Paper';
import Typography from 'material-ui/Typography';
import PropTypes from 'prop-types';
import Genetics from './Genetics.jsx';
import GeneticsMap from './GeneticsMap.jsx';

class Species extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      name: props.name,
      genetics: { traits:[] },
    };
  }

  componentDidMount() {
    axios.get(`/species/${this.props.name}`)
      .then((res) => {
        const s = res.data;
        this.setState({ genetics:s });
      });
  }

  render() {
    return (
      <div>
        <Paper elevation={4}>
          <Typography type="headline" component="h2">
            {this.state.name}
          </Typography>
          <Genetics traits={this.state.genetics.traits}/>
          <GeneticsMap />
        </Paper>
      </div>
    );
  }


}

Species.propTypes = {
  name: PropTypes.string.isRequired,
};

module.exports = Species;
