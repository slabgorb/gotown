import { Avatar, Card, CardContent, CardHeader, Chip, CircularProgress, Typography } from '@material-ui/core';
import { withStyles } from '@material-ui/core/styles';
import fetch from 'fetch-hoc';
import inflection from 'inflection';
import PropTypes from 'prop-types';
import React from 'react';


const _ = require('underscore');

const styles = theme => ({
  root: {
    minWidth: 300,
    margin: theme.spacing.unit,
  },
  avatar: {
    backgroundColor: theme.palette.primary.main,
  },
});

const CultureCard = ({
  data, classes, loading, error,
}) => {
  if (loading) {
    return (<CircularProgress className={classes.progress} />);
  }
  if (Object.keys(error).length) {
    return (<div>{error.message}</div>);
  }
  const { name, marital_strategies, namer_names } = data;
  return (
    <Card className={classes.root}>
      <CardHeader
        avatar={(<Avatar area-label="Culture" className={classes.avatar}>C</Avatar>)}
        title={inflection.titleize(name)}
      />
      <CardContent>
        <Typography>Marital Strategies</Typography>
        {marital_strategies.map(m => <Chip label={inflection.titleize(m)} />)}
        <Typography>Namers</Typography>
        {_.map(namer_names, (v, k) => <Chip label={inflection.titleize(v)} avatar={<Avatar>{k[0].toUpperCase()}</Avatar>} />)}
      </CardContent>
    </Card>
  );
};

CultureCard.propTypes = {
  data: PropTypes.object,
  classes: PropTypes.object.isRequired,
  loading: PropTypes.bool,
  error: PropTypes.object,
};

CultureCard.defaultProps = {
  error: {},
  data: {},
  loading: true,
};

module.exports = fetch(({ id }) => `/api/cultures/${id}`)(withStyles(styles)(CultureCard));
