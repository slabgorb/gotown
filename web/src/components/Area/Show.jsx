import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';
import areaApi from './api';

const styles = () => ({
  root: {},
});

class AreaShow extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      name: props.match.params.name,
    };
  }

  componentWillMount() {
    return areaApi.get(this.state.name).then((data) => {
      this.setState({
        name: data.name,
      });
    });
  }

  render() {
    const { classes } = this.props;
    return (<div className={ classes.root }>{this.state.name}</div>);
  }
}

AreaShow.propTypes = {
  match: PropTypes.object.isRequired,
  classes: PropTypes.object.isRequired,
};


export default withStyles(styles)(AreaShow);