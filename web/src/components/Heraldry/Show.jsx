import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';

const styles = {};

const Show = ({ src, width }) => (
  <img width={width} alt="heraldry" src={src} />
);

Show.propTypes = {
  src: PropTypes.string.isRequired,
  width: PropTypes.number,
};

Show.defaultProps = {
  width: 270,
};

export default withStyles(styles)(Show);
