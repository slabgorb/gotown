import { Typography } from '@material-ui/core';
import { withStyles } from '@material-ui/core/styles';
import PropTypes from 'prop-types';
import React from 'react';

const styles = theme => ({
  root: {
    margin: [0, theme.spacing.unit],
  },
});

class Trait extends React.Component {
  constructor(props) {
    super(props);
    this.state = { expanded: false };
    this.handleExpandClick = this.handleExpandClick.bind(this);
  }

  handleExpandClick() {
    this.setState({ expanded: !this.state.expanded });
  }

  render() {
    const { name, classes, value } = this.props;
    return (
      <div className={classes.root}>
        <Typography variant="caption">
          {name}
        </Typography>
        <Typography variant="subheading" component="h3">
          {value}
        </Typography>
      </div>
    );
  }
}

Trait.propTypes = {
  name: PropTypes.string.isRequired,
  classes: PropTypes.object.isRequired,
  value: PropTypes.string.isRequired,
};

export default withStyles(styles)(Trait);
