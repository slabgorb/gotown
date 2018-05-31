import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';

import speciesApi from './api';
import { MenuList } from '../App';

const styles = () => ({
  heading: {
    fontSize: 14,
  },
});

class SpeciesList extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      list: props.list,
    };
    //  this.handleClick = this.handleClick.bind(this);
  }

  componentWillMount() { 
    if (this.state.list.length > 0) { return; }
    speciesApi.getAll().then(data => this.setState({ list: data }));
  }

  render() {
    const { classes, handleClick } = this.props;
    if (this.state.list.length === 0) {
      return null;
    }
    return (<MenuList
      heading="Species"
      classes={classes}
      list={this.state.list}
      handleClick={v => handleClick(`species/${v}`)()}
    />);
  }
}

SpeciesList.propTypes = {
  handleClick: PropTypes.func.isRequired,
  classes: PropTypes.object.isRequired,
  list: PropTypes.array,
};

SpeciesList.defaultProps = {
  list: [],
};

export default withStyles(styles)(SpeciesList);
