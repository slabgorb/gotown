import React from 'react';
import PropTypes from 'prop-types';
import Chip from 'material-ui/Chip';
import { withStyles } from 'material-ui/styles';

const _ = require('underscore');

const styles = {
  patternSet: {
    margin: '5px 0px',
  }
};

const PatternChipper = ({ pattern, classes }) => {
  const re = new RegExp('{{\.([A-Za-z]+)}}', 'g');
  const results = [];
  let m;
  do {
    m = re.exec(pattern);
    if (m) {
      results.push(m[1]);
    }
  } while (m);

  return (<div className={classes.patternSet} key={pattern}>{_.map(results, r => (<Chip key={r} label={r} />))}</div>);
};

PatternChipper.propTypes = {
  pattern: PropTypes.string.isRequired,
  classes: PropTypes.object.isRequired,
};


module.exports = withStyles(styles)(PatternChipper);