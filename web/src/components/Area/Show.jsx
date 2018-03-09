import React from 'react';
import { withStyles } from 'material-ui/styles';
import areaApi from './api';

const styles = () => ({

});

class Show extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      name: '',
    };
  }

  componentWillMount() {
    return areaApi.get(this.state.name).then((data) => {
      this.setState({
        name: data.name,
      });
    });
  }
}

export default withStyles(styles)(Show)