import React from 'react';
import PropTypes from 'prop-types';
import Paper from 'material-ui/Paper';
import Typography from 'material-ui/Typography';
import { withStyles } from 'material-ui/styles';
import inflection from 'inflection';
import wordsApi from './api';
import NameList from '../NameList';

const _ = require("underscore");

const styles = theme => ({
  cardContent: {
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
  get({ name }) {
    wordsApi.get(name)
      .then((s) => {
        this.setState({
          name: s.name,
          dictionary: s.dictionary,
          loaded: true,
        });
      });
  }

  render() {
    const { classes } = this.props;
    if (!this.state.loaded) {
      return (<div>loading</div>);
    }
    return (
      <Paper>
        <Typography variant="headline" component="h1" className={classes.headline}>
          {inflection.titleize(this.state.name)}
        </Typography>
        {
          _.map(_.filter(this.state.dictionary, d => d.length > 0), (v, k) => (<NameList key={k} title={k} listItems={v} />))
        }
      </Paper>
    )
  }
} 

Words.propTypes = {
  match: PropTypes.object.isRequired,
  classes: PropTypes.object.isRequired,
};
module.exports = withStyles(styles)(Words);
