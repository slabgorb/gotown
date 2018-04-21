import React from 'react';
import { withStyles } from 'material-ui/styles';
import PropTypes from 'prop-types';
import Tabs, { Tab } from 'material-ui/Tabs';
import AppBar from 'material-ui/AppBar';
import inflection from 'inflection';

const _ = require('underscore');

const style = () => ({
  root: {
    marginTop: 64,
  }
});

class TabBar extends React.Component {
  constructor(props) {
    super(props);
    this.state = ({ value: 1 });
    this.handleChange = this.handleChange.bind(this);
  }

  handleChange(event, value) {
    const { onChange } = this.props;
    this.setState({ value });
    onChange(value);
  }

  render() {
    const { classes, tabs } = this.props;
    const { value } = this.state;
    return (
      <AppBar position="static">
        <Tabs value={value} classes={{ root: classes.root }} onChange={this.handleChange}>
          { _.map(tabs, t => (<Tab label={inflection.humanize(t)} key={t} />)) }
        </Tabs>
      </AppBar>
    );
  }
}

TabBar.propTypes = {
  tabs: PropTypes.array.isRequired,
  classes: PropTypes.object.isRequired,
  onChange: PropTypes.func.isRequired,
}

module.exports = withStyles(style)(TabBar);
