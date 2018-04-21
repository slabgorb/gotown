import React from 'react';
import { withStyles } from 'material-ui/styles';
import PropTypes from 'prop-types';
import AppBar from 'material-ui/AppBar';
import { ChromosomeShow } from '../Chromosome';
import Genetics from './Genetics';
import GeneticsMap from './GeneticsMap';
import speciesApi from './api';
import { PageTitle, TabBar } from '../App';

const styles = () => ({
  tabs: {
    marginTop: 72,
  },
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
    const { value, name } = this.state;
    const { classes } = this.props;
    return (
      <div>
        <PageTitle title={name} capitalize subtitle="Species" />
        <AppBar position="static" color="primary">
          <TabBar onChange={this.handleChange} tabs={['example', 'expression', 'map']} />
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
  classes: PropTypes.object.isRequired,
};

module.exports = withStyles(styles)(Species);
