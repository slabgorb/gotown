import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';
import areaApi from './api';
import { MenuList } from '../App';

const styles = () => ({
});

class AreaList extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      list: [],
    };
    this.handleDelete = this.handleDelete.bind(this);
  }

  componentWillMount() {
    areaApi.getAll().then(data => this.setState({ list: data }));
  }


  handleDelete(event, item) {
    event.stopPropagation();
    areaApi.delete(item).then(() =>
      areaApi.getAll().then(data => this.setState({ list: data })));
  }


  render() {
    const { list } = this.state;
    if (list.length === 0) {
      return null;
    }
    const { classes, handleClick } = this.props;
    return (
      <MenuList
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
