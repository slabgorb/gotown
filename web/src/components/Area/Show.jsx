import React from 'react';
import { withStyles } from 'material-ui/styles';
import areaApi from './api';

const styles = () => ({

});

class AreaShow extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      name: this.props.name,
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
    return (<div>{this.state.name}</div>);
  }
}

export default withStyles(styles)(AreaShow);
