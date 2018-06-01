import Button from '@material-ui/core/Button';
import Dialog from '@material-ui/core/Dialog';
import DialogTitle from '@material-ui/core/DialogTitle';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import { withStyles } from '@material-ui/core/styles';
import PropTypes from 'prop-types';
import React from 'react';


const style = theme => ({
  root: {
    transition: '0.3s all',
    backgroundColor: theme.palette.primary.main,
    color: 'white',
  },
});

const SimpleDialog = ({
  classes,
  open,
  title,
  message,
  onYes,
  onCancel,
}) => (
  <Dialog open={open} className={classes.root}>
    <DialogTitle>{title}</DialogTitle>
    <Typography>{message}</Typography>
    <Toolbar>
      <Button onClick={onYes}>Yes</Button>
      <Button onClick={onCancel}>Cancel</Button>
    </Toolbar>
  </Dialog>
);

SimpleDialog.propTypes = {
  classes: PropTypes.object.isRequired,
  open: PropTypes.bool.isRequired,
  title: PropTypes.string.isRequired,
  message: PropTypes.string.isRequired,
  onYes: PropTypes.func.isRequired,
  onCancel: PropTypes.func.isRequired,
};

module.exports = withStyles(style)(SimpleDialog);
