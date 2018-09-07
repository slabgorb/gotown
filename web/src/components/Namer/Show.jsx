import { withStyles } from '@material-ui/core/styles';
import PropTypes from 'prop-types';
import React from 'react';
import { PageTitle, TabBar } from '../App';
import { WordsShow } from '../Words';
import namerApi from './api';
import PatternChipper from './PatternChipper';
import Random from './Random';
import { Grid, Card, CardContent, CardActionArea, Typography, CardHeader } from '@material-ui/core';

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
  leftGrid: {
    
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
      value: 1,
      id: 0,
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
    namerApi.get(id)
      .then((s) => {
        this.setState({
          name: s.name,
          id: s.id,
          wordsID: s.words_id,
          patterns: s.patterns,
          loaded: true,
        });
      });
  }


  render() {
    const { classes, showAppBar } = this.props;
    const {
      loaded,
      patterns,
      name,
      id,
      wordsID,
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
        <Grid container>
          <Grid item xs={6}>
            <Card>
              <CardContent>
                <Random id={id} />
                { patternChips }
              </CardContent>
            </Card>
          </Grid>
          <Grid item xs={6}>
            <Card>
              <CardHeader title="Words" />
              <CardContent>
                <WordsShow showAppBar={false} match={{ params: { name: wordsID } }} />
              </CardContent>
            </Card>
          </Grid>
        </Grid>
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
