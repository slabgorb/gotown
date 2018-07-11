import { withStyles } from '@material-ui/core/styles';
import PropTypes from 'prop-types';
import React from 'react';
import MenuList from '../App/MenuList';
import WithLoading from '../App/WithLoading';
import areaApi from './api';

const styles = () => ({
});

const MenuListWithLoading = WithLoading(MenuList);

class AreaList extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      list: [],
      isLoading: true,
    };
    this.handleDelete = this.handleDelete.bind(this);
  }

  componentWillMount() {
    areaApi.getAll().then(data => this.setState({ list: data, isLoading: false }));
  }


  handleDelete(event, item) {
    event.stopPropagation();
    areaApi.delete(item).then(() =>
      areaApi.getAll().then(data => this.setState({ list: data })));
  }


  render() {
    const { list, isLoading } = this.state;
    const { classes, handleClick } = this.props;
    return (
      <MenuListWithLoading
        isLoading={isLoading}
        deletable
        classes={classes}
        heading="Towns"
        list={list}
        handleClick={v => handleClick(`towns/${v}`)()}
        handleDelete={this.handleDelete}
      />
    );
  }
}

AreaList.propTypes = {
  handleClick: PropTypes.func.isRequired,
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(AreaList);
