import { withStyles } from 'material-ui/styles';
import PropTypes from 'prop-types';
import React from 'react';
import { HeraldryShow } from '../Heraldry';
import areaApi from './api';

const styles = () => ({
  root: {},
});

class AreaShow extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      id: props.match.params.id,
      name: '',
      image: '',
    };
  }

  componentWillMount() {
    return areaApi.get(this.state.id).then((data) => {
      this.setState({
        name: data.name,
        image: data.image,
      });
    });
  }

  render() {
    const { classes } = this.props;
    const { name, image } = this.state;
    return (
      <div>
        <div className={classes.root}>{name}</div>
        <HeraldryShow src={image} size={270} />
      </div>
    );
  }
}

AreaShow.propTypes = {
  match: PropTypes.object.isRequired,
  classes: PropTypes.object.isRequired,
};


export default withStyles(styles)(AreaShow);
