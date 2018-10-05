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
    const { loading, classes } = props;
    if (!loading) return (<Component {...props} />);
    return (<CircularProgress className={classes.progress} />);
  };
  wrapped.propTypes = {
    loading: PropTypes.bool,
    classes: PropTypes.object.isRequired,
  };
  wrapped.defaultProps = {
    loading: true,
  };
  return withStyles(styles)(wrapped);
};

export default WithLoading;
