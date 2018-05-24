import { withStyles } from 'material-ui/styles';
import PropTypes from 'prop-types';
import React from 'react';
import { PageTitle, TabBar } from '../App';
import { WordsShow } from '../Words';
import PatternChipper from './PatternChipper';
import Random from './Random';
import namerApi from './api';

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
  tabs: { marginTop: 72 },
});

class Namer extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      name: props.match.params.name,
      wordsName: '',
      patterns: [],
      loaded: false,
      value: 1,
      id: 0,
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
    namerApi.get(name)
      .then((s) => {
        this.setState({
          name: s.name,
          id: s.id,
          wordsName: s.words,
          patterns: s.patterns,
          loaded: true,
        });
      });
  }

  handleChange(value) {
    this.setState({ value });
  }

  render() {
    const { classes, showAppBar } = this.props;
    const {
      loaded,
      patterns,
      name,
      value,
      wordsName,
      id,
    } = this.state;
    if (!loaded) {
      return (<div>loading</div>);
    }
    const patternChips = [];
    _.each(patterns, (p, i) => {
      patternChips.push((<PatternChipper key={i} pattern={p} />));
    });
    return (
      <div>
        { showAppBar && (<PageTitle title={name} classes={classes} titleize subtitle="Namer" />) }
        <TabBar onChange={this.handleChange} tabs={['patterns', 'words', 'test']} />
        { value === 0 && patternChips}
        { value === 1 &&
          (<WordsShow showAppBar={false} match={{ params: { name: wordsName } }} />)
        }
        { value === 2 && (<Random id={id} />)}
      </div>
    );
  }
}

Namer.propTypes = {
  match: PropTypes.object.isRequired,
  classes: PropTypes.object.isRequired,
  showAppBar: PropTypes.bool,
};

Namer.defaultProps = {
  showAppBar: true,
};

module.exports = withStyles(styles)(Namer);
