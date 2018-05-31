import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';

const styles = theme => ({
  cardContent: {
    backgroundColor: theme.palette.background.paper,
  },
  cardHeader: {
    fontFamily: 'Raleway',
    fontSize: '14',
    subheader: {
      fontFamily: 'Raleway',
    },
  },
  list: {
  },
  listItem: {
    fontFamily: 'Raleway',
    fontSize: '12',
  },
});

class Dictionary extends React.Component {

}


module.exports = withStyles(styles)(Dictionary);