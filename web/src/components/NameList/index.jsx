import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import { withStyles } from '@material-ui/core/styles';
import PropTypes from 'prop-types';
import React from 'react';

const _ = require('underscore');

const styles = theme => ({
  root: {
  },
  cardContent: {
    backgroundColor: theme.palette.background.paper,
  },
  cardHeader: {
    fontSize: '14',
  },
  list: {
  },
  listItem: {
    fontSize: '12',
  },
});

const NameList = ({
  classes,
  listItems,
}) => (
    <List className={classes.list} dense disablePadding>
      { _.uniq(listItems).map(f =>
        (<ListItem className={classes.listItem} key={f}>{f}</ListItem>))}
    </List>
);

NameList.propTypes = {
  listItems: PropTypes.array.isRequired,
  classes: PropTypes.object.isRequired,
};
NameList.defaultProps = {
};


module.exports = withStyles(styles)(NameList);
