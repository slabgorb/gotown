import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';
import { MenuList } from '../App';
import namerApi from './api';

const styles = () => ({
  heading: {
    fontSize: 14,
  },
});
class NamersList extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      list: props.list,
    };
  }

  componentWillMount() {
    if (this.state.list.length > 0) { return; }
    namerApi.getAll().then(data => this.setState({ list: data }));
  }

  render() {
    if (this.state.list.length === 0) {
      return null;
    }
    const { classes, handleClick } = this.props;
    return (<MenuList
      heading="Namers"
      classes={classes}
      list={this.state.list}
      handleClick={v => handleClick(`namers/${v}`)()}
    />);
  }
}

NamersList.propTypes = {
  handleClick: PropTypes.func.isRequired,
  classes: PropTypes.object.isRequired,
  list: PropTypes.array,
};

NamersList.defaultProps = {
  list: [],
};

export default withStyles(styles)(NamersList);
