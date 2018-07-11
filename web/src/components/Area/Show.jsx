import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import CardHeader from '@material-ui/core/CardHeader';
import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import { withStyles } from '@material-ui/core/styles';
import inflection from 'inflection';
import PropTypes from 'prop-types';
import React from 'react';
import { PageTitle, TabBar } from '../App';
import { BeingList } from '../Being';
import { BarChart, RadarChart } from '../Charts/';
import { HeraldryShow } from '../Heraldry';
import areaApi from './api';

const _ = require('underscore');

const styles = () => ({
  card: {
    maxWidth: 300,
    minWidth: 300,
    margin: 20,
  },
  charts: {
    display: 'flex',
  },
  root: {
    flexGrow: 1,
  },
  paper: {},
});

class AreaShow extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      id: props.match.params.id,
      name: '',
      image: '',
      residents: [],
      tab: 0,
    };
    this.handleChange = this.handleChange.bind(this);
  }

  componentWillMount() {
    return areaApi.get(this.state.id).then((data) => {
      this.setState({
        name: data.name,
        image: data.image,
        residents: data.residents.beings,
        icon: data.icon,
      });
    });
  }


  handleChange(tab) {
    this.setState({ tab });
  }

  render() {
    const { classes } = this.props;
    const {
      name,
      image,
      residents,
      tab,
      icon,
    } = this.state;
    /* eslint-disable no-param-reassign */
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
    /* eslint-enable no-param-reassign */

    const histoData = _.map(
      residents.reduce(histogramReducer, {}),
      (value, title) => ({ title, value }),
    );
    const traits = [
      'height',
      'weight',
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

    const radarDataSets = traits.map(t => [
      { axes: _.map(residents.reduce(traitReducer(t), {}), (value, axis) => ({ axis, value })) },
    ]);
    const radarCharts = radarDataSets.map((ds, i) =>
      (
        <Grid item xs={12} sm={6} key={traits[i]}>
          <Card className={classes.card}>
            <CardHeader title={inflection.titleize(traits[i])} />
            <CardContent>
              <RadarChart data={ds} w={240} h={240} />
            </CardContent>
          </Card>
        </Grid>
      ));

    const tab1 = (<div><HeraldryShow src={image} size={270} /><BarChart data={histoData} /></div>);
    const tab2 = (<div className={classes.root}><Grid container spacing={8}>{radarCharts}</Grid></div>);
    const tab3 = (<BeingList beings={residents} />);

    return (
      <div className={classes.root}>
        <PageTitle title={name} titleize subtitle="Area" icon={(<img src={icon} width={32} height={32} alt="icon" />)} />
        <TabBar value={tab} onChange={this.handleChange} tabs={['details', 'traits', 'list']} />
        <Paper className={classes.paper} >
          { tab === 0 && tab1 }
          { tab === 1 && tab2 }
          { tab === 2 && tab3 }
        </Paper>
      </div>
    );
  }
}

AreaShow.propTypes = {
  match: PropTypes.object.isRequired,
  classes: PropTypes.object.isRequired,
};


export default withStyles(styles)(AreaShow);
