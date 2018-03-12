import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';
import { withRouter } from 'react-router-dom';
import List, { ListItem, ListItemText } from 'material-ui/List';
import inflection from 'inflection';
import areaApi from './api';

const styles = () => {

};

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


  render() {
    if (this.state.list.length === 0) {
      return null;
    }
    return (
      <List component="nav">
        {this.state.list.map(item => (
          <ListItem button divider key={item} onClick={() => this.handleClick(item)}>
            <ListItemText primary={inflection.titleize(item)} />
          </ListItem>
          ))}
      </List>
    );
  }
}

AreaList.propTypes = {
  history: PropTypes.object.isRequired,

};

export default withRouter(withStyles(styles)(AreaList));
