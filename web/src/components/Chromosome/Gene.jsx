import React from 'react';
import PropTypes from 'prop-types';
import TextField from 'material-ui/TextField';
import { withStyles } from 'material-ui/styles';

const styles = () => {

};

const Gene = ({ value, onChange }) =>
  (
    <div>
      <TextField
        maxLength={8}
        value={value}
        onChange={onChange} 
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
    </div>
  );

Gene.propTypes = {
  value: PropTypes.string,
  onChange: PropTypes.func.isRequired,
};

Gene.defaultProps = {
  value: 'ffffff',
};

export default withStyles(styles)(Gene);
