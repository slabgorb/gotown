import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';
import { withRouter } from 'react-router-dom';
import areaApi from './api';
import { MenuList } from '../App';

const styles = () => ({
  deleteButton: {},
});

class AreaList extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      list: [],
    };
    this.handleClick = this.handleClick.bind(this);
  }

  componentWillMount() {
    areaApi.getAll().then(data => this.setState({ list: data }));
  }

  handleClick(value) {
    this.props.history.push(`/towns/${value}`);
  }

  handleDelete(event, item) {
    event.stopPropagation();
    areaApi.delete(item).then(() =>
      areaApi.getAll().then(data => this.setState({ list: data })));
  }


  render() {
    const { classes } = this.props;
    const { list } = this.state;
    if (list.length === 0) {
      return null;
    }
    return (
      <MenuList
        deletable
        classes={classes}
        heading="Towns"
        list={list}
        handleClick={this.handleClick}
        handleDelete={this.handleDelete}
      />
    );
  }
}

AreaList.propTypes = {
  history: PropTypes.object.isRequired,
  classes: PropTypes.object.isRequired,
};

export default withRouter(withStyles(styles)(AreaList));
