import React from 'react';
import Paper from 'material-ui/Paper';
import Typography from 'material-ui/Typography';
import PropTypes from 'prop-types';
import Drawer from 'material-ui/Drawer';
import Button from 'material-ui/Button';
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
      chromosomeDrawerOpen: false,
    };
    this.get = this.get.bind(this);
    this.toggleChromosomeDrawer = this.toggleChromosomeDrawer.bind(this);
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

  toggleChromosomeDrawer(chromosomeDrawerOpen) {
    return () => this.setState({ chromosomeDrawerOpen });
  }

  render() {
    if (this.state.genetics.traits.length === 0) {
      return (<div />);
    }
    return (
      <div>
        <Drawer
          anchor="top"
          open={this.state.chromosomeDrawerOpen}
          onClose={this.toggleChromosomeDrawer(false)}
          variant="persistant"
        >
          <Button onClick={this.toggleChromosomeDrawer(false)}>Close</Button>
          <ChromosomeShow speciesName={this.state.name} />
        </Drawer>
        <Paper elevation={4}>
          <Typography variant="headline" component="h1">
            {this.state.name}
          </Typography>
          <Button onClick={this.toggleChromosomeDrawer(true)}>Chromosome Test</Button>
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
