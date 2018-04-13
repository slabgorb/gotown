import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';
import { withRouter } from 'react-router-dom';
import List, { ListItem, ListItemText } from 'material-ui/List';
import inflection from 'inflection';
import wordsApi from './api';

const styles = () => {

};

class WordsList extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      list: [],
    };
    this.handleClick = this.handleClick.bind(this);
  }

  componentWillMount() {
    wordsApi.getAll().then(data => this.setState({ list: data }));
  }

  handleClick(value) {
    this.props.history.push(`/words/${value}`);
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

WordsList.propTypes = {
  history: PropTypes.object.isRequired,

};

export default withRouter(withStyles(styles)(WordsList));
