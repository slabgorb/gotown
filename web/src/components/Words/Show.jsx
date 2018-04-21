import React from 'react';
import PropTypes from 'prop-types';
import Paper from 'material-ui/Paper';
import Typography from 'material-ui/Typography';
import { withStyles } from 'material-ui/styles';
import wordsApi from './api';
import NameList from '../NameList';
import { PageTitle, TabBar } from '../App';

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
      dictionary: {},
      loaded: false,
      value: 0,
    };
    this.get = this.get.bind(this);
    this.getDictByValue = this.getDictByValue.bind(this);
    this.getKeyByValue = this.getKeyByValue.bind(this);
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

  getKeyByValue(value) {
    const keys = _.keys(this.state.dictionary);
    return keys[value];
  }

  getDictByValue(value) {
    return this.state.dictionary[this.getKeyByValue[value]];
  }

  handleChange(value) {
    this.setState({ value });
  }

  get({ name }) {
    wordsApi.get(name)
      .then((s) => {
        this.setState({
          name: s.name,
          dictionary: _.pick(s.dictionary, d => d.length > 0),
          loaded: true,
        });
      });
  }

  render() {
    const { classes } = this.props;
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
        <PageTitle title={name} titleize /> 
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
                )
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
};
module.exports = withStyles(styles)(Words);
