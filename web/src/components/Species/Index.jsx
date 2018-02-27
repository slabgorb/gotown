import 'typeface-raleway';
import React from 'react';
import axios from 'axios';
import Paper from 'material-ui/Paper';
import Typography from 'material-ui/Typography';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';
import List, { ListItem, ListItemText } from 'material-ui/List';

const styles = () => {

};

class Index extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      list: [],
    };
  }
  componentDidMount() {
    axios.get('/species')
      .then((res) => {
        const s = res.data;
        this.setState({ list: s });
      });
  }
  render() {
    if (this.state.list.length == 0) {
      return null;
    }
    return (
      <List component="nav">
        {this.state.list.map((item) => {
          return (
            <ListItem button key={item}>
              <ListItemText primary={item} />
            </ListItem>
          );
        })}
      </List>
    );
  }
}

export default withStyles(styles)(Index);