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
      namerShown: 0,
      loaded: false,
      maritalStrategies: [],
      value: 0,
    };
    this.get = this.get.bind(this);
    this.handleChange = this.handleChange.bind(this);
    this.handleNamerClick = this.handleNamerClick.bind(this);
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

  handleNamerClick(namerShown) {
    this.setState({ namerShown });
  }

  render() {
    const { classes } = this.props;
    const { value, namers, namerShown } = this.state;
    if (!this.state.loaded) {
      return (<div>loading</div>);
    }
    return (
      <div className={classes.root}>
        <PageTitle title={this.state.name} titleize subtitle="Culture" />
        <TabBar value={value} onChange={this.handleChange} tabs={['names', 'marriage']} />
        { value === 0 && (
          <div>
            <TabBar value={namerShown} onChange={this.handleNamerClick} tabs={_.pluck(_.values(namers), 'name')} />
            <NamersShow showAppBar={false} match={{ params: { name: _.values(namers)[namerShown].name } }} />
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
