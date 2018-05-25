import Chip from 'material-ui/Chip';
import { withStyles } from 'material-ui/styles';
import PropTypes from 'prop-types';
import React from 'react';

const _ = require('underscore');

const styles = {
  patternSet: {
    margin: '5px 0px',
  },
};


const PatternChipper = ({ pattern, classes }) => {
  const re = /({{\.[A-Za-z]+}})/g;
  let parts = pattern.split(re);
  let m;
  const subRe = /{{\.([A-Za-z]+)}}/;
  parts = _.map(parts, (p) => {
    m = subRe.exec(p);
    if (m) {
      return (<Chip key={m[1]} label={m[1]} />);
    }
    return (<span>{p}</span>);
  });
  return (
    <div className={classes.patternSet} key={pattern}>{parts}</div>
  );
};

PatternChipper.propTypes = {
  pattern: PropTypes.string.isRequired,
  classes: PropTypes.object.isRequired,
};


module.exports = withStyles(styles)(PatternChipper);
