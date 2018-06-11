import Paper from '@material-ui/core/Paper';
import { withStyles } from '@material-ui/core/styles';
import PropTypes from 'prop-types';
import React from 'react';
import { PageTitle, TabBar } from '../App';
import NameList from '../NameList';
import wordsApi from './api';

const _ = require('underscore');

const styles = theme => ({
  root: {
    backgroundColor: theme.palette.background.paper,
  },
  cardHeader: {
    fontFamily: 'Raleway',
    fontSize: '14',
    subheader: {
      fontFamily: 'Raleway',
    },
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
    wordsApi.get(backupName)
      .then((s) => {
        const dictionary = Object.assign(s.dictionary, this.state.dictionary);
        this.setState({ dictionary });
      });
  }

  handleChange(value) {
    this.setState({ value });
  }

  get({ id }) {
    wordsApi.get(id)
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
      value,
      name,
      dictionary,
    } = this.state;
    if (!loaded) {
      return (<div>loading</div>);
    }
    return (
      <div>
        { showAppBar && (<PageTitle title={name} titleize />) }
        <TabBar onChange={this.handleChange} tabs={_.keys(dictionary)} />
        <Paper classes={{ root: classes.root }}>
          { _.map(_.keys(dictionary), (k, i) => {
              if (i === value) {
                return (
                  <NameList
                    key={k}
                    title={k}
                    listItems={dictionary[k]}
                  />
                );
              }
              return '';
            })
          }
        </Paper>
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
