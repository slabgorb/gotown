import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';
import { withRouter } from 'react-router-dom';
import List, { ListItem, ListItemText } from 'material-ui/List';
import IconButton from 'material-ui/IconButton';
import DeleteForeverIcon from 'material-ui-icons/DeleteForever';
import inflection from 'inflection';
import areaApi from './api';

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
    if (this.state.list.length === 0) {
      return null;
    }
    return (
      <List component="nav">
        {this.state.list.map(item => (
          <ListItem button divider key={item} onClick={() => this.handleClick(item)}>
            <ListItemText primary={inflection.titleize(item)} />
            <IconButton className={classes.deleteButton} onClick={e => this.handleDelete(e, item)}>
              <DeleteForeverIcon />
            </IconButton>
          </ListItem>
          ))}
      </List>
    );
  }
}

AreaList.propTypes = {
  history: PropTypes.object.isRequired,
  classes: PropTypes.object.isRequired,
};

export default withRouter(withStyles(styles)(AreaList));
