import { withStyles } from '@material-ui/core/styles';
import PropTypes from 'prop-types';
import React from 'react';
import { PageTitle, TabBar } from '../App';

import wordsApi from './api';
import { Typography } from '@material-ui/core';
import Dictionary from './Dictionary';

const _ = require('underscore');

const styles = theme => ({
  root: {
    backgroundColor: theme.palette.background.paper,
  },
  list: {
  },
  listItem: {
    fontFamily: 'Raleway',
    fontSize: '12',
  },
});

class Words extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      name: props.match.params.name,
      backupName: props.backupName,
      dictionary: {},
      loaded: false,
      value: 0,
    };
    this.get = this.get.bind(this);
    this.getDictByValue = this.getDictByValue.bind(this);
    this.getKeyByValue = this.getKeyByValue.bind(this);
    this.handleChange = this.handleChange.bind(this);
    this.getBackup = this.getBackup.bind(this);
  }

  componentDidMount() {
    this.get(this.props.match.params);
  }

  componentWillReceiveProps(nextProps) {
    if (this.props.match.params.name !== nextProps.match.params.name) {
      this.get(nextProps.match.params);
    }
  }

  getKeyByValue(value) {
    const keys = _.keys(this.state.dictionary);
    return keys[value];
  }

  getDictByValue(value) {
    return this.state.dictionary[this.getKeyByValue[value]];
  }


  getBackup({ backupName }) {
    console.log({ backupName })
    wordsApi.get(backupName)
      .then((s) => {
        const dictionary = Object.assign(s.dictionary, this.state.dictionary);
        this.setState({ dictionary });
      });
  }

  handleChange(value) {
    this.setState({ value });
  }

  get({ name, id }) {
    wordsApi.get(name, id)
      .then((s) => {
        this.setState({
          name: s.name,
          backupName: s.backup,
          dictionary: _.pick(s.dictionary, d => d.length > 0),
          loaded: true,
        });
      })
      .then(() =>
        this.state.backupName !== '' && this.getBackup(this.state))
      .then(() => this.setState({ loaded: true }));
  }

  render() {
    const { classes, showAppBar } = this.props;
    const {
      loaded,
      name,
      dictionary,
    } = this.state;
    if (!loaded) {
      return (<div>loading</div>);
    }
    return (
      <div>
        { showAppBar && (<PageTitle title={name} titleize />) }
        { _.map(_.keys(dictionary), k => (<Dictionary k={k} classes={classes} key={k} dictionary={dictionary[k]}/>)) }
      </div>
    );
  }
}



Words.propTypes = {
  match: PropTypes.object.isRequired,
  classes: PropTypes.object.isRequired,
  showAppBar: PropTypes.bool,
  backupName: PropTypes.string,
};

Words.defaultProps = {
  showAppBar: true,
  backupName: '',
};
module.exports = withStyles(styles)(Words);
