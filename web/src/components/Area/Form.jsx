import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';
import Radio, { RadioGroup } from 'material-ui/Radio';
import { FormLabel, FormControl, FormControlLabel } from 'material-ui/Form';
import inflection from 'inflection';
import areaApi from './api';

const styles = () => ({
  root: {
    padding: '15',
  },
  formControl: {},
});


class Form extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      cultures: [],
      species: [],
      loaded: false,
    };
  }

  componentWillMount() {
    Promise.all([
      areaApi.getCultures(),
      areaApi.getSpecies(),
    ]).then(([cultures, species]) =>
      this.setState({
        cultures,
        species,
        loaded: true,
        currentCulture: '',
        currentSpecies: '',
      }));
  }

  handleChange(control) {
    return (event, value) => {
      this.setState({
        [control]: value,
      });
    };
  }

  radioGroup(legend, name, list, value, onChange) {
    const { classes } = this.props;
    return (
      <FormControl component="fieldset" required className={classes.formControl}>
        <FormLabel component="legend">{legend}</FormLabel>
        <RadioGroup name={name} value={value} onChange={onChange}>
          {list.map(f => (
            <FormControlLabel key={f} value={f} control={<Radio />} label={inflection.titleize(f)} />
       ))}
        </RadioGroup>
      </FormControl>
    );
  }

  render() {
    const { classes } = this.props;
    if (!this.state.loaded) {
      return (<div>Loading</div>);
    }
    return (
      <div className={classes.root}>
        {this.radioGroup('Species', 'species', this.state.species, this.state.currentSpecies, this.handleChange('currentSpecies'))}
        {this.radioGroup('Culture', 'culture', this.state.cultures, this.state.currentCulture, this.handleChange('currentCulture'))}
      </div>
    );
  }
}

Form.propTypes = {
  classes: PropTypes.object.isRequired,
};

module.exports = withStyles(styles)(Form);
