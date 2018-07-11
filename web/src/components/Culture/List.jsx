import { withStyles } from '@material-ui/core/styles';
import PropTypes from 'prop-types';
import React from 'react';
import { withRouter } from 'react-router-dom';
import MenuList from '../App/MenuList';
import WithLoading from '../App/WithLoading';
import cultureApi from './api';


const MenuListWithLoading = WithLoading(MenuList);
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
      isLoading: true,
    };
  }

  componentWillMount() {
    cultureApi.getAll().then(data => this.setState({ list: data, isLoading: false }));
  }

  render() {
    const { list, isLoading } = this.state;
    const { classes, handleClick } = this.props;
    return (<MenuListWithLoading
      isLoading={isLoading}
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
