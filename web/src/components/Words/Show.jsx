import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';

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

class Words extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      name: props.match.params.name,
      namers: [],
      loaded: false,
      maritalStrategies: [],
    };
    this.get = this.get.bind(this);
  }
} 

Culture.propTypes = {
  match: PropTypes.object.isRequired,
  classes: PropTypes.object.isRequired,
};
module.exports = withStyles(styles)(Words);