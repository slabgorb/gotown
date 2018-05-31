import React from 'react';
import PropTypes from 'prop-types';
import TextField from '@material-ui/core/TextField';
import { withStyles } from '@material-ui/core/styles';

const styles = () => {

};

const Gene = ({ value, onChange, hasFocus }) =>
  (
    <TextField
      maxLength={8}
      value={value}
      onChange={onChange}
      autoFocus={hasFocus}
      label={(
        <span
          style={
          {
            display: 'inline-block',
            backgroundColor: `#${value}`,
            width: '20px',
            height: '20px',
          }}
        />
      )}
    />
  );

Gene.propTypes = {
  value: PropTypes.string,
  onChange: PropTypes.func.isRequired,
  hasFocus: PropTypes.bool.isRequired,
};

Gene.defaultProps = {
  value: 'ffffff',
};

export default withStyles(styles)(Gene);
