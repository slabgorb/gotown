import { withStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import PropTypes from 'prop-types';
import React from 'react';

const styles = () => ({
  root: {},
  header: {
    padding: '24px 24px',
  },
  paragraph: {
    padding: '16px 24px',
  },
});

const Info = ({ classes, header, paras }) => (
  <div className={classes.root}>
    <Typography variant="headline" className={classes.header}>
      {header}
    </Typography>
    {paras.map(p => (
      <Typography component="p" className={classes.paragraph}>
        {p}
      </Typography>
    ))}
  </div>
);

Info.propTypes = {
  classes: PropTypes.object.isRequired,
  header: PropTypes.string.isRequired,
  paras: PropTypes.array.isRequired,
};

module.exports = withStyles(styles)(Info)