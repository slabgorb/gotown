import { Avatar, Card, CardContent, CardHeader, CircularProgress } from '@material-ui/core';
import { withStyles } from '@material-ui/core/styles';
import fetch from 'fetch-hoc';
import inflection from 'inflection';
import PropTypes from 'prop-types';
import React from 'react';
import { BarChart } from '../Charts';

const _ = require('underscore');

const styles = theme => ({
  root: {
    minWidth: 300,
  },
  avatar: {
    backgroundColor: theme.palette.primary.main,
  },
});

const SpeciesCard = ({ data, classes, loading, error }) => {
  if (loading) {
    return (<CircularProgress className={classes.progress} />);
  }
  
  if (error) {
    return (<div>{error}</div>);
  }
  const { name, demography } = data;
  const chartData = _.map(demography, d => ({ value: d.cumulative_percent, title: `${d.max_age}` }));
  return (
    <Card className={classes.root}>
      <CardHeader
        avatar={(<Avatar area-label="Species" className={classes.avatar}>S</Avatar>)}
        title={inflection.titleize(name)}
      />
      <CardContent>
        <BarChart
          data={chartData}
        />
      </CardContent>
    </Card>
  );
};

SpeciesCard.propTypes = {
  data: PropTypes.object.isRequired,
  classes: PropTypes.object.isRequired,
  loading: PropTypes.bool.isRequired,
  error: PropTypes.string.isRequired,
};

// SpeciesCard.defaultProps = {
//   error: '',
// };

module.exports = fetch(({ id }) => `/api/species/${id}`)(withStyles(styles)(SpeciesCard));