import React from 'react';
import { withStyles } from 'material-ui/styles';
import PropTypes from 'prop-types';
import inflection from 'inflection';
import cultureApi from './api';
import { NamersList, NamersShow } from '../Namer';
import { PageTitle, TabBar } from '../App';

const _ = require('underscore');

const styles = theme => ({
  root: {
    backgroundColor: theme.palette.background.paper,
    fontFamily: 'Montserrat',
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

  handleChange(value) {
    this.setState({ value });
  }

  render() {
    const { classes } = this.props;
    const { value, namers } = this.state;
    if (!this.state.loaded) {
      return (<div>loading</div>);
    }

    const tabs = ['marriage'];
    tabs.push(..._.keys(namers).map(s => `${s} names`));
    
    return (
      <div className={classes.root}>
        <PageTitle title={this.state.name} titleize subtitle="Culture" />
        <TabBar value={value} onChange={this.handleChange} tabs={tabs} />
        { value > 0 && (
          <NamersShow showAppBar={false} match={{ params: { name: _.values(namers)[value - 1].name } }} />
        )}
        { value === 0 && (
          <div className="flex-container">
            { _.map(this.state.maritalStrategies, v => (<span key={v}>{inflection.titleize(v)}</span>)) }
          </div>
        )}
      </div>
    );
  }
}

Culture.propTypes = {
  match: PropTypes.object.isRequired,
  classes: PropTypes.object.isRequired,
};

module.exports = withStyles(styles)(Culture);
