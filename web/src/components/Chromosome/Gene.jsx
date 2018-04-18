import React from 'react';
import PropTypes from 'prop-types';
import TextField from 'material-ui/TextField';
import { withStyles } from 'material-ui/styles';

const styles = () => {

};

const Gene = ({value, onChange}) =>
  (
    <TextField defaultValue={value} onChange={onChange} />
  );

Gene.propTypes = {
  value: PropTypes.string,
  onChange: PropTypes.func.isRequired,
};

Gene.defaultProps = {
  value: 'ffffff',
};

export default withStyles(styles)(Gene);