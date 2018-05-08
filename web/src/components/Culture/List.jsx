import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';
import { withRouter } from 'react-router-dom';
import { MenuList } from '../App';
import cultureApi from './api';

const styles = () => ({
  heading: {
    fontSize: 14,
  },
});

class CultureList extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      list: [],
    };
  }

  componentWillMount() {
    cultureApi.getAll().then(data => this.setState({ list: data }));
  }

  render() {
    const { list } = this.state;
    if (list.length === 0) {
      return null;
    }
    const { classes, handleClick } = this.props;
    return (<MenuList
      heading="Cultures"
      classes={classes}
      list={list}
      handleClick={v => handleClick(`cultures/${v}`)()}
    />);
  }
}

CultureList.propTypes = {
  handleClick: PropTypes.func.isRequired,
  classes: PropTypes.object.isRequired,

};

export default withRouter(withStyles(styles)(CultureList));
