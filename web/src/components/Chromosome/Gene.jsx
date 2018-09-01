import React from 'react';
import PropTypes from 'prop-types';
import TextField from '@material-ui/core/TextField';
import { withStyles } from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';

const styles = () => ({
  colorSquare: {
    width: '20px',
    height: '20px',
    marginTop: '15px',
  },
  input: {
    width: "6em",
  },
  root: {
    marginLeft:10,
  }
});

const Gene = ({ value, onChange, hasFocus, classes }) =>
  (
    <Grid container spacing={32} className={classes.root}>
      <Grid item xs={1} className={classes.colorSquare} style={({backgroundColor: `#${value}`})}/>
      <Grid item xs={2}>
        <TextField
          className={classes.input}
          maxLength={8}
          value={value}
          onChange={onChange}
          autoFocus={hasFocus}
        />
      </Grid>
    </Grid>
  );

Gene.propTypes = {
  value: PropTypes.string,
  classes: PropTypes.object.isRequired,
  onChange: PropTypes.func.isRequired,
  hasFocus: PropTypes.bool.isRequired,
};

Gene.defaultProps = {
  value: 'ffffff',
};

export default withStyles(styles)(Gene);
