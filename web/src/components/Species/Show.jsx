import React from 'react';
import { withStyles } from 'material-ui/styles';
import Typography from 'material-ui/Typography';
import PropTypes from 'prop-types';
import Tabs, { Tab } from 'material-ui/Tabs';
import AppBar from 'material-ui/AppBar';
import { ChromosomeShow } from '../Chromosome';
import Genetics from './Genetics';
import GeneticsMap from './GeneticsMap';
import speciesApi from './api';

const styles = () => ({
  tabs: {
    
  }
});

class Species extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      name: props.match.params.name,
      genetics: { traits: [] },
      value: 1,
    };
    this.get = this.get.bind(this);
    this.handleChange = this.handleChange.bind(this);
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


  handleChange(event, value) {
    this.setState({ value });
  }


  render() {
    if (this.state.genetics.traits.length === 0) {
      return (<div />);
    }
    const { value } = this.state;
    return (
      <div>
        <Typography variant="headline" component="h1">
          {this.state.name}
        </Typography>
        <AppBar position="static" color="default">
          <Tabs onChange={this.handleChange} centered>
            <Tab label="example" />
            <Tab label="expression" />
            <Tab label="map" />
          </Tabs>
        </AppBar>
        { value === 0 && (<ChromosomeShow speciesName={this.state.name} />) }
        { value === 1 && (<Genetics traits={this.state.genetics.traits} />) }
        { value === 2 && (<GeneticsMap traits={this.state.genetics.traits} />) }
      </div>
    );
  }
}

Species.propTypes = {
  match: PropTypes.object.isRequired,
  // classes: PropTypes.object.isRequired,
};

module.exports = withStyles(styles)(Species);
