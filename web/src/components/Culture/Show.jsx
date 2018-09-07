import { withStyles } from '@material-ui/core/styles';
import inflection from 'inflection';
import PropTypes from 'prop-types';
import React from 'react';
import { PageTitle, TabBar } from '../App';
import { Grid, Paper } from '@material-ui/core';
import { NamersShow } from '../Namer';
import cultureApi from './api';

const _ = require('underscore');

const styles = theme => ({
  root: {
    backgroundColor: theme.palette.background.paper,
    fontFamily: 'Vollkorn',
  },
});

class Culture extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      name: props.match.params.name,
      namers: [],
      loaded: false,
      maritalStrategies: [],
      value: 0,
    };
    this.get = this.get.bind(this);
  }

  componentDidMount() {
    this.get(this.props.match.params);
  }

  componentWillReceiveProps(nextProps) {
    if (this.props.match.params.name !== nextProps.match.params.name) {
      this.get(nextProps.match.params);
    }
  }

  get({ id }) {
    cultureApi.get(id)
      .then((s) => {
        this.setState({
          name: s.name,
          namers: s.namers,
          maritalStrategies: s.marital_strategies,
          loaded: true,
        });
      });
  }

  render() {
    const { classes } = this.props;
    const { namers } = this.state;
    if (!this.state.loaded) {
      return (<div>loading</div>);
    }
    return (
      <div className={classes.root}>
        <PageTitle title={this.state.name} titleize subtitle="Culture" />
        <Grid container>
          { _.map(namers, v => (
          <Grid item xs={6} key={v.name}>
            <Paper>
              <NamersShow
                showAppBar={false}
                match={{ params: { id: v.id, name: v.name } }}
              />
            </Paper>

          </Grid>
          ))}

          <Grid item xs={6}>
          { _.map(
                this.state.maritalStrategies,
                v => (<span key={v}>{inflection.titleize(v)}</span>),
            )}
          </Grid>
        </Grid>
      </div>
    );
  }
}

Culture.propTypes = {
  match: PropTypes.object.isRequired,
  classes: PropTypes.object.isRequired,
};

module.exports = withStyles(styles)(Culture);
