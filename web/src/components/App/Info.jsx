import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import { withStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import PropTypes from 'prop-types';
import React from 'react';

const styles = theme => ({
  root: {},
  paper: {
    margin: [theme.spacing.unit, theme.spacing.unit * 2].join(' '),
  },
  header: {
    padding: theme.spacing.unit * 2,
  },
  paragraph: {
    padding: theme.spacing.unit * 2,
  },
  children: {
    padding: theme.spacing.unit * 2,
  },
});

const Info = ({
  classes,
  header,
  paras,
  xs,
  sm,
  md,
  children,
}) => (
  <Grid item className={classes.root} xs={xs} sm={sm} md={md}>
    <Paper className={classes.paper}>
      <Typography variant="headline" className={classes.header}>
        {header}
      </Typography>
      <div className={classes.children}>
        {children}
      </div>
      {paras.map(p => (
        <Typography component="p" className={classes.paragraph} key={p}>
          {p}
        </Typography>
      ))}
    </Paper>
  </Grid>
);

Info.propTypes = {
  classes: PropTypes.object.isRequired,
  header: PropTypes.string.isRequired,
  paras: PropTypes.array.isRequired,
  xs: PropTypes.number,
  sm: PropTypes.number,
  md: PropTypes.number,
  children: PropTypes.node,
};

Info.defaultProps = {
  xs: 12,
  sm: 6,
  md: 3,
  children: null,
};
module.exports = withStyles(styles)(Info)