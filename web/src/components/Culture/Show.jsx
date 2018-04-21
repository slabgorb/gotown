import React from 'react';
import { withStyles } from 'material-ui/styles';
import PropTypes from 'prop-types';
import Tabs, { Tab } from 'material-ui/Tabs';
import AppBar from 'material-ui/AppBar';
import inflection from 'inflection';
import cultureApi from './api';
import { NamersList } from '../Namer';
import { PageTitle } from '../App';

const _ = require('underscore');

const styles = theme => ({
  root: {
    backgroundColor: theme.palette.background.paper,
    fontFamily: 'Montserrat',
  },
  headline: {
    fontFamily: 'Montserrat',
    marginLeft: '20',
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
    this.handleChange = this.handleChange.bind(this);
  }

  componentDidMount() {
    this.get(this.props.match.params);
  }

  componentWillReceiveProps(nextProps) {
    if (this.props.match.params.name !== nextProps.match.params.name) {
      this.get(nextProps.match.params);
    }
  }

  get({ name }) {
    cultureApi.get(name)
      .then((s) => {
        this.setState({
          name: s.name,
          namers: s.namers,
          maritalStrategies: s.marital_strategies,
          loaded: true,
        });
      });
  }

  handleChange(event, value) {
    this.setState({ value });
  }

  render() {
    const { classes } = this.props;
    const { value } = this.state;
    if (!this.state.loaded) {
      return (<div>loading</div>);
    }
    return (
      <div>
        <PageTitle title={this.state.name} titleize subtitle="Culture" />
        <AppBar position="static" color="primary">
          <Tabs onChange={this.handleChange} centered>
            <Tab label="names" />
            <Tab label="marriage" />
          </Tabs>
        </AppBar>
        { value === 0 && (
          <div className="flex-container">
            <NamersList list={_.pluck(_.values(this.state.namers), 'name')} />
          </div>) }
        { value === 1 && (
          <div className="flex-container">
            { _.map(this.state.maritalStrategies, v => (<span key={v}>{inflection.titleize(v)}</span>)) }
          </div>) }
      </div>
    );
  }
}

Culture.propTypes = {
  match: PropTypes.object.isRequired,
  classes: PropTypes.object.isRequired,
};

module.exports = withStyles(styles)(Culture);
