import { withStyles } from '@material-ui/core/styles';
import PropTypes from 'prop-types';
import React from 'react';
import { PageTitle, TabBar } from '../App';
import { ChromosomeShow } from '../Chromosome';
import Genetics from './Genetics';
import GeneticsMap from './GeneticsMap';
import speciesApi from './api';

const styles = () => ({
  root: {},
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


  handleChange(value) {
    this.setState({ value });
  }


  render() {
    const { genetics, value, name } = this.state;
    if (genetics.traits.length === 0) {
      return (<div />);
    }
    const { classes } = this.props;
    return (
      <div className={classes.root}>
        <PageTitle title={name} capitalize subtitle="Species" />
        <TabBar onChange={this.handleChange} tabs={['example', 'expression', 'map']} />
        { value === 0 && (<ChromosomeShow speciesName={name} />) }
        { value === 1 && (<Genetics traits={genetics.traits} />) }
        { value === 2 && (<GeneticsMap traits={genetics.traits} />) }
      </div>
    );
  }
}

Species.propTypes = {
  match: PropTypes.object.isRequired,
  classes: PropTypes.object.isRequired,
};

module.exports = withStyles(styles)(Species);
