import { GridList, GridListTile } from '@material-ui/core';
import { withStyles } from '@material-ui/core/styles';
import PropTypes from 'prop-types';
import React from 'react';
import Trait from './Trait';

const _ = require('underscore');

const styles = theme => ({
  root: {
    margin: theme.spacing.unit,
    minWidth: 200,
  },
  listItem: {
    root: {

    },

  },
});

const findTrait = (traits, name) => _.find(traits, e => e.name === name);

const Expression = ({ expressionMap, traits, classes }) => (
  <GridList cols={2} cellHeight={50} className={classes.root}>
    { _.map(expressionMap, (v, k) => (
      <GridListTile key={k}>
        <Trait {...findTrait(traits, k)} value={v} />
      </GridListTile>)) }
  </GridList>
);

Expression.propTypes = {
  expressionMap: PropTypes.object.isRequired,
  traits: PropTypes.array.isRequired,
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(Expression);
