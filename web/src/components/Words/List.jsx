import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import { MenuList } from '../App';
import wordsApi from './api';

const styles = () => ({
  heading: {
    fontSize: 14,
  },
});

class WordsList extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      list: [],
    };
  }

  componentWillMount() {
    wordsApi.getAll().then(data => this.setState({ list: data }));
  }

  render() {
    if (this.state.list.length === 0) {
      return null;
    }
    const { classes, handleClick } = this.props;
    return (<MenuList
      heading="Words"
      classes={classes}
      list={this.state.list}
      handleClick={v => handleClick(`words/${v}`)()}
    />);
  }
}

WordsList.propTypes = {
  handleClick: PropTypes.func.isRequired,
  classes: PropTypes.object.isRequired,

};

export default withStyles(styles)(WordsList);
