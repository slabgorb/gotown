import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';
import { withRouter } from 'react-router-dom';
import List, { ListItem, ListItemText } from 'material-ui/List';
import speciesApi from '../../api/speciesApi';

const styles = () => {

};

class SpeciesList extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      list: [],
    };
    this.handleClick = this.handleClick.bind(this);
  }

  componentWillMount() { speciesApi.getAll().then(data => this.setState({ list: data })); }

  handleClick(value) {
    this.props.history.push(`/species/${value}`);
  }


  render() {
    if (this.state.list.length === 0) {
      return null;
    }
    return (
      <List component="nav">
        {this.state.list.map(item => (
          <ListItem button divider key={item} onClick={() => this.handleClick(item)}>
            <ListItemText primary={item} />
          </ListItem>
          ))}
      </List>
    );
  }
}

SpeciesList.propTypes = {
  history: PropTypes.object.isRequired,

};

export default withRouter(withStyles(styles)(SpeciesList));
