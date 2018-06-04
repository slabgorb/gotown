import { withStyles } from '@material-ui/core/styles';
import PropTypes from 'prop-types';
import React from 'react';
import { BeingList } from '../Being';
import { Histogram } from '../Charts';
import { HeraldryShow } from '../Heraldry';
import areaApi from './api';

const _ = require('underscore');

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
      residents: [],
    };
  }

  componentWillMount() {
    return areaApi.get(this.state.id).then((data) => {
      this.setState({
        name: data.name,
        image: data.image,
        residents: data.residents.beings,
      });
    });
  }

  render() {
    const { classes } = this.props;
    const { name, image, residents } = this.state;

    const histogramReducer = (memo, current) => {
      if (current.age in memo) {
        memo[current.age] += 1;
      } else {
        memo[current.age] = 1;
      }
      return memo;
    };

    const histoData = _.map(residents.reduce(histogramReducer, {}), (value, title) => ({ title, value }));
    return (
      <div>
        <div className={classes.root}>{name}</div>
        <Histogram data={histoData} />
        <HeraldryShow src={image} size={270} />
        <BeingList beings={residents} />
      </div>
    );
  }
}

AreaShow.propTypes = {
  match: PropTypes.object.isRequired,
  classes: PropTypes.object.isRequired,
};


export default withStyles(styles)(AreaShow);
