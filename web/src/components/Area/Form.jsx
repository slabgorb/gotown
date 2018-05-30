import inflection from 'inflection';
import AutoRenewIcon from 'material-ui-icons/Autorenew';
import Button from 'material-ui/Button';
import { FormControl, FormControlLabel, FormLabel } from 'material-ui/Form';
import Grid from 'material-ui/Grid';
import IconButton from 'material-ui/IconButton';
import { MenuItem } from 'material-ui/Menu';
import Radio, { RadioGroup } from 'material-ui/Radio';
import Select from 'material-ui/Select';
import TextField from 'material-ui/TextField';
import { withStyles } from 'material-ui/styles';
import PropTypes from 'prop-types';
import React from 'react';
import { PageTitle } from '../App';
import List from './List';
import areaApi from './api';

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
      currentSize: 8,
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
      size: this.state.currentSize,
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
              key={f.id}
              value={f.name}
              control={<Radio />}
              label={inflection.titleize(f.name)}
            />
       ))}
        </RadioGroup>
      </FormControl>
    );
  }

  render() {
    const { classes } = this.props;
    const {
      name,
      currentSize,
      loaded,
      species,
      cultures,
      currentCulture,
      currentSpecies,
    } = this.state;
    if (!loaded) {
      return (<div>Loading</div>);
    }
    return (
      <form className={classes.root} onSubmit={this.submitForm}>
        <PageTitle title="Create Town" />
        <List />
        <Grid container>
          <Grid item xs={6}>
            <div className="flex-container">
              <TextField id="name-tf" label="Name" placeholder="Name" value={name} className={classes.textField} onChange={this.handleChange('name')} />
              <IconButton className={classes.avatar} onClick={this.clickRandomTownName}>
                <AutoRenewIcon />
              </IconButton>
            </div>
          </Grid>
          <Grid item xs={6}>
            <Button variant="raised" color="primary" type="submit" disabled={!(currentCulture && currentSpecies)} className={classes.button}>Create</Button>
          </Grid>
          <Grid item xs={6}>
            {this.radioGroup('Species', 'species', species, currentSpecies, this.handleChange('currentSpecies'))}
          </Grid>
          <Grid item xs={6}>
            {this.radioGroup('Culture', 'culture', cultures, currentCulture, this.handleChange('currentCulture'))}
          </Grid>
          <Grid item xs={6}>
            <FormControl classname={classes.formControl}>
              <Select
                value={currentSize}
                onChange={this.handleChange('currentSize')}
                name="size"
              >
                <MenuItem value={1}>Hut </MenuItem>
                <MenuItem value={2}>Cottage</MenuItem>
                <MenuItem value={3}>House</MenuItem>
                <MenuItem value={4}>Tower</MenuItem>
                <MenuItem value={5}>Castle</MenuItem>
                <MenuItem value={6}>Hamlet</MenuItem>
                <MenuItem value={7}>Palace</MenuItem>
                <MenuItem value={8}>Village</MenuItem>
                <MenuItem value={9}>Town</MenuItem>
                <MenuItem value={10}>City</MenuItem>
                <MenuItem value={11}>Region</MenuItem>
                <MenuItem value={12}>NationState</MenuItem>
                <MenuItem value={13}>Empire</MenuItem>
              </Select>
            </FormControl>

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
