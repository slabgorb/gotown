import React from 'react';
import PropTypes from 'prop-types';
import Paper from 'material-ui/Paper';
import { withStyles } from 'material-ui/styles';
import Typography from 'material-ui/Typography';
import inflection from 'inflection';
import namerApi from './api';
import { WordsShow } from '../Words';
import NameList from '../NameList';
import PatternChipper from './PatternChipper';

const _ = require('underscore');

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

class Namer extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      name: props.match.params.name,
      wordsName: '',
      patterns: [],
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
    namerApi.get(name)
      .then((s) => {
        this.setState({
          name: s.name,
          wordsName: s.words,
          patterns: s.patterns,
          loaded: true,
        });
      });
  }


  render() {
    const { classes } = this.props;
    if (!this.state.loaded) {
      return (<div>loading</div>);
    }
    const patternChips = []
    _.each(this.state.patterns, (p, i) => {
      patternChips.push((<PatternChipper key={i} pattern={p} />));
    });
    return (
      <Paper elevation={4} className={classes.root}>;
        <Typography variant="headline" component="h1" className={classes.headline}>
          {inflection.titleize(this.state.name)}
        </Typography>
        {patternChips}
        <WordsShow match={{ params: { name: this.state.wordsName } }} />
      </Paper>
    );
  }
}

Namer.propTypes = {
  match: PropTypes.object.isRequired,
  classes: PropTypes.object.isRequired,
};


module.exports = withStyles(styles)(Namer);