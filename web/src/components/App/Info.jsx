import { Card, Grid, Typography } from '@material-ui/core';
import { withStyles } from '@material-ui/core/styles';
import PropTypes from 'prop-types';
import React from 'react';

const styles = theme => ({
  root: {},
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
    <Card className={classes.card}>
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
    </Card>
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