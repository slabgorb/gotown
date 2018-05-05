import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';
import Radio, { RadioGroup } from 'material-ui/Radio';
import { FormLabel, FormControl, FormControlLabel } from 'material-ui/Form';
import Button from 'material-ui/Button';
import Grid from 'material-ui/Grid';
import IconButton from 'material-ui/IconButton';
import TextField from 'material-ui/TextField';
import AutoRenewIcon from 'material-ui-icons/Autorenew';
import inflection from 'inflection';
import areaApi from './api';
import { PageTitle } from '../App';
import List from './List';

const styles = theme => ({
  root: {
    padding: '15',
  },
  formControl: {},
  textField: {},
  avatar: {
  },
  button: {
    margin: theme.spacing.unit,
  },
});


class Form extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      cultures: [],
      species: [],
      loaded: false,
      name: '',
      currentCulture: null,
      currentSpecies: null,
    };
    this.clickRandomTownName = this.clickRandomTownName.bind(this);
    this.submitForm = this.submitForm.bind(this);
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
    return (event) => {
      this.setState({
        [control]: event.target.value,
      });
    };
  }

  submitForm(event) {
    event.preventDefault();
    const params = {
      culture: this.state.currentCulture,
      species: this.state.currentSpecies,
      name: this.state.name,
    };
    areaApi.create(params).then(data => console.log(data));
  }

  clickRandomTownName() {
    areaApi.name().then((name) => {
      this.setState({
        name,
      });
    });
  }

  radioGroup(legend, name, list, value, onChange) {
    const { classes } = this.props;
    return (
      <FormControl component="fieldset" required className={classes.formControl}>
        <FormLabel component="legend">{legend}</FormLabel>
        <RadioGroup name={name} value={value} onChange={onChange}>
          {list.map(f => (
            <FormControlLabel
              key={f}
              value={f}
              control={<Radio />}
              label={inflection.titleize(f)}
            />
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
      <form className={classes.root} onSubmit={this.submitForm}>
        <PageTitle title="Create Town" />
        <List />
        <Grid container>
          <Grid item xs={6}>
            <div className="flex-container">
              <TextField id="name-tf" label="Name" placeholder="Name" value={this.state.name} className={classes.textField} onChange={this.handleChange('name')} />
              <IconButton className={classes.avatar} onClick={this.clickRandomTownName}>
                <AutoRenewIcon />
              </IconButton>
            </div>
          </Grid>
          <Grid item xs={6}>
            <Button variant="raised" color="primary" type="submit" disabled={!(this.state.currentCulture && this.state.currentSpecies)} className={classes.button}>Create</Button>
          </Grid>
          <Grid item xs={6}>
            {this.radioGroup('Species', 'species', this.state.species, this.state.currentSpecies, this.handleChange('currentSpecies'))}
          </Grid>
          <Grid item xs={6}>
            {this.radioGroup('Culture', 'culture', this.state.cultures, this.state.currentCulture, this.handleChange('currentCulture'))}
          </Grid>

        </Grid>
      </form>
    );
  }
}

Form.propTypes = {
  classes: PropTypes.object.isRequired,
};

module.exports = withStyles(styles)(Form);
