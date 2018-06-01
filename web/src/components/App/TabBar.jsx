import React from 'react';
import { withStyles } from '@material-ui/core/styles';
import PropTypes from 'prop-types';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import AppBar from '@material-ui/core/AppBar';
import inflection from 'inflection';

const _ = require('underscore');

const style = () => ({
  root: {
  }
});

class TabBar extends React.Component {
  constructor(props) {
    super(props);
    this.state = ({ value: props.value });
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
  value: PropTypes.number,
  tabs: PropTypes.array.isRequired,
  classes: PropTypes.object.isRequired,
  onChange: PropTypes.func.isRequired,
};

TabBar.defaultProps = {
  value: 1,
};

module.exports = withStyles(style)(TabBar);
