import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';
import BottomNavigation, { BottomNavigationAction } from 'material-ui/BottomNavigation';
import FingerprintIcon from 'material-ui-icons/Fingerprint';
import LanguageIcon from 'material-ui-icons/Language';
import PlaceIcon from 'material-ui-icons/Place';

const styles = () => {
  
};

class SimpleBottomNavigation extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      value: 0,
    };
    this.handleChange = this.handleChange.bind(this);
  }

  handleChange(event, value) {
    this.setState({ value });
  };

  render() {
    const { classes } = this.props;
    const { value } = this.state;

    return (
      <BottomNavigation
        value={value}
        onChange={this.handleChange}
        showLabels
        className={classes.root}
      >
        <BottomNavigationAction label="Species" icon={<FingerprintIcon />} />
        <BottomNavigationAction label="Cultures" icon={<LanguageIcon />} />
        <BottomNavigationAction label="Towns" icon={<PlaceIcon />} />
      </BottomNavigation>
    );
  }
}

SimpleBottomNavigation.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(SimpleBottomNavigation);
