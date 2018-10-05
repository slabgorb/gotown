import { Typography } from '@material-ui/core';
import Grid from '@material-ui/core/Grid';
import { withStyles } from '@material-ui/core/styles';
import PropTypes from 'prop-types';
import React from 'react';
import { CultureCard } from '../Culture';
import cultureApi from '../Culture/api';
import { SpeciesCard } from '../Species';
import speciesApi from '../Species/api';
import PageTitle from './PageTitle';

const _ = require('underscore');

const style = theme => ({
  info: {
    margin: theme.spacing.unit * 2,
  },
  root: {
    marginLeft: theme.spacing.unit * 4,
  },
});


class Home extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      speciesList: [],
      cultureList: [],
    };
  }


  componentWillMount() {
    this.getSpecies();
    this.getCultures();
  }

  getSpecies() {
    if (this.state.speciesList.length > 0) { return; }
    speciesApi.getAll().then(data => this.setState({ speciesList: data }));
  }

  getCultures() {
    if (this.state.cultureList.length > 0) { return; }
    cultureApi.getAll().then(data => this.setState({ cultureList: data }));
  }

  render() {
    const { classes } = this.props;
    const { speciesList, cultureList } = this.state;
    return (
      <div className={classes.root}>
        <PageTitle title="Home" titleize classes={classes} />
        <Typography variant="display1">Species</Typography>
        <Grid container className={classes.info}>
          { speciesList.map(s => (<SpeciesCard key={s.id} id={s.id} />))}
        </Grid>
        <Typography variant="display1">Cultures</Typography>
        <Grid container className={classes.info}>
          { cultureList.map(s => (<CultureCard key={s.id} id={s.id} />))}
        </Grid>
      </div>
    );
  }
}

Home.propTypes = {
  classes: PropTypes.object.isRequired,
};

module.exports = withStyles(style)(Home);
