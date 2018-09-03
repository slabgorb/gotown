import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import FormLabel from '@material-ui/core/FormLabel';
import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import Typography from '@material-ui/core/Typography';
import { withStyles } from '@material-ui/core/styles';
import AutoRenewIcon from '@material-ui/icons/Autorenew';
import PropTypes from 'prop-types';
import React from 'react';
import speciesApi from '../Species/api';
import chromosomeApi from './api';
import Expression from './Expression';
import Gene from './Gene'; 

const _ = require('underscore');

const styles = theme => ({
  root: {
    flexGrow: 1,
    padding: theme.spacing.unit * 4,
  },
  avatar: {
    paddingRight: theme.spacing.unit * 3,
  },
  chromosomeFull: {
    fontSize: 14,
    whiteSpace: 'pre-wrap',
  },
});

class Show extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      genes: [],
      changing: false,
      expressionMap: {},
      focusedGene: -1,
    };
    this.changeGene = this.changeGene.bind(this);
    this.clickRandomChromosome = this.clickRandomChromosome.bind(this);
    this.getExpression = this.getExpression.bind(this);
  }

  componentDidMount() {
    this.clickRandomChromosome()
  }

  getExpression() {
    const { speciesID } = this.props;
    return speciesApi.getExpression(speciesID, this.state.genes)
      .then(expressionMap => this.setState({ expressionMap }));
  }

  changeGene(index) {
    if (this.state.changing) {
      return () => { };
    }
    const f = (event) => {
      const { value } = event.target;
      this.setState(prevState => ({
        genes: prevState.genes.map((g, i) => {
          if (i === index) { return value; }
          return g;
        }),
      }), this.getExpression);
    };
    return f.bind(this);
  }

  clickGeneInMap(focusedGene) {
    return () => {
      this.setState({ focusedGene });
    };
  }

  clickRandomChromosome() {
    this.setState({
      changing: true,
    });
    chromosomeApi.random().then(({ genes }) => {
      this.setState({
        genes,
      });
    }).then(this.getExpression)
      .then(() => this.setState({ changing: false }));
  }
  render() {
    const { classes, traits } = this.props;
    return (
      <Paper>
        
        <Grid
          container
          spacing={24}
          className={classes.root}
        >  
          <Grid container item>
            <Grid item xs={2} sm={1}>
              <Button
                variant="extendedFab"
                aria-label="randomize"
                className={classes.avatar}
                onClick={this.clickRandomChromosome}
                onKeyUp={this.checkHex}
                color="primary"
              >
                <AutoRenewIcon />
                Randomize
              </Button>
            </Grid>

          </Grid>
          <Grid container>
            <Grid item xs={4} sm={3}>
              <FormControl>
                <FormLabel label="check genetics" />
                {_.map(this.state.genes, (g, i) => (
                  <Gene
                    value={g}
                    key={i}
                    hasFocus={i === this.state.focusedGene}
                    onChange={this.changeGene(i)}
                  />
                ))}
              </FormControl>
            </Grid>
            <Grid item xs={8} sm={9}>
              <Expression expressionMap={this.state.expressionMap} traits={traits} />
            </Grid>
          </Grid>

        </Grid >
      </Paper>

    );
  }
}

Show.propTypes = {
  classes: PropTypes.object.isRequired,
  speciesName: PropTypes.string.isRequired,
  speciesID: PropTypes.string.isRequired,
  traits: PropTypes.array.isRequired,
};

export default withStyles(styles)(Show);
