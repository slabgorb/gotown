import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import CardHeader from '@material-ui/core/CardHeader';
import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import Typography from '@material-ui/core/Typography';
import { withStyles } from '@material-ui/core/styles';
import inflection from 'inflection';
import PropTypes from 'prop-types';
import React from 'react';
import { BeingList } from '../Being';
import { BarChart, RadarChart } from '../Charts/';
import { HeraldryShow } from '../Heraldry';
import areaApi from './api';

const _ = require('underscore');

const styles = () => ({
  card: {
    maxWidth: 350,
    minWidth: 350,
    margin: 20,
  },
  charts: {
    display: 'flex',
  },
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
    const traitReducer = trait => (memo, current) => {
      const ec = current.expression[trait];
      if (ec in memo) {
        memo[ec] += 1;
      } else {
        memo[ec] = 1;
      }
      return memo;
    };

    const histoData = _.map(residents.reduce(histogramReducer, {}), (value, title) => ({ title, value }));
    const traits = [
      'agreeableness',
      'conscientiousness',
      'ear shape',
      'emotionality',
      'extraversion',
      'eye color',
      'eye shape',
      'eyebrows',
      'face shape',
      'hair color',
      'honesty/humility',
      'lip shape',
      'openness',
      'skin color',
    ];

    const radarDataSets = traits.map(t => [{ axes: _.map(residents.reduce(traitReducer(t), {}), (value, axis) => ({ axis, value })) }]);
    const radarCharts = radarDataSets.map((ds, i) =>
      (
        <Grid item xs={12} sm={6}>
          <Card className={classes.card}>
            <CardHeader title={inflection.titleize(traits[i])} />
            <CardContent>
              <RadarChart data={ds} w={300} h={300} />
            </CardContent>
          </Card>
        </Grid>
      ));
    return (
      <Paper className={classes.root} >
        <Typography variant="display1">{name}</Typography>
        <Grid container spacing={8}>
          {radarCharts}
        </Grid>
        <BarChart data={histoData} />
        <HeraldryShow src={image} size={270} />
        <BeingList beings={residents} />
      </Paper>
    );
  }
}

AreaShow.propTypes = {
  match: PropTypes.object.isRequired,
  classes: PropTypes.object.isRequired,
};


export default withStyles(styles)(AreaShow);
