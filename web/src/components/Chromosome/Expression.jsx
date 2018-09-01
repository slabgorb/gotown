import PropTypes from 'prop-types';
import React from 'react';
import Trait from './Trait';
import Grid from '@material-ui/core/Grid';

const _ = require('underscore');

const findTrait = (traits, name) => _.find(traits, (e) => e.name === name)

const Expression = ({ expressionMap, traits }) => (
  <Grid container xs={12} sm={12} lg={12} spacing={16}>
    { _.map(expressionMap, (v, k) => ( 
      <Grid xs={6} item>
        <Trait {...findTrait(traits, k)} value={v}/>
      </Grid> )) }
  </Grid>
)

Expression.propTypes = {
  expressionMap: PropTypes.object.isRequired,
  traits: PropTypes.array.isRequired,
};

export default Expression;
