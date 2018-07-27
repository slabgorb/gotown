import CircularProgress from '@material-ui/core/CircularProgress';
import { withStyles } from '@material-ui/core/styles';
import PropTypes from 'prop-types';
import React from 'react';

const styles = theme => ({
  progress: {
    margin: theme.spacing.unit * 2,
  },
});

const WithLoading = (Component) => {
  const wrapped = (props) => {
    const { isLoading, classes } = props;
    if (!isLoading) return (<Component { ...props } />);
    return (<CircularProgress className={classes.progress} />);
  };
  wrapped.propTypes = {
    isLoading: PropTypes.bool,
    classes: PropTypes.object.isRequired,
  };
  wrapped.defaultProps = {
    isLoading: true,
  };
  return withStyles(styles)(wrapped);
};

export default WithLoading;
