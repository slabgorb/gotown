import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';
import IconButton from 'material-ui/IconButton';
import AutoRenewIcon from 'material-ui-icons/Autorenew';
import { FormLabel, FormControl } from 'material-ui/Form';
import Grid from 'material-ui/Grid';
import Gene from './Gene';
import chromosomeApi from './api';
import speciesApi from '../Species/api';
import Expression from './Expression';

const _ = require('underscore');

const styles = () => ({
  root: {
    flexGrow: 1,
  },
  avatar: {},
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
    };
    this.changeGene = this.changeGene.bind(this);
    this.clickRandomChromosome = this.clickRandomChromosome.bind(this);
    this.getExpression = this.getExpression.bind(this);
  }

  getExpression() {
    const { speciesName } = this.props;
    return speciesApi.getExpression(speciesName, this.state.genes)
      .then(expressionMap => this.setState({ expressionMap }));
  }

  changeGene(index) {
    if (this.state.changing) {
      return () => {};
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
    const { classes } = this.props;
    return (

      <Grid
        container
        spacing={24}
        className={classes.root}
        justify="center"
      >
        <Grid item xs={12}>
          <Grid container>
            <Grid item xs={2} sm={1}>
              <IconButton
                className={classes.avatar}
                onClick={this.clickRandomChromosome}
                onKeyUp={this.checkHex}
              >
                <AutoRenewIcon />
              </IconButton>
            </Grid>
            <Grid item xs={10} sm={8} className={classes.chromosomeFull}>
              <div className="gene-table">
                { this.state.genes.length > 0
                  ? this.state.genes.map((g, i) => (
                    <div
                      className="gene"
                      style={{
                        borderBottomColor: `#${g}`,
                        borderBottomWidth: '2px',
                        borderBottomStyle: 'solid',
                      }}
                      key={g}
                    >
                      {g}
                      {i % 4 === 0 ? (<br />) : '' }
                    </div>))
                  : ''
                }
              </div>
            </Grid>
          </Grid>
        </Grid>
        <Grid item xs={12} sm={2}>
          <FormControl>
            <FormLabel label="check genetics" />
            { _.map(this.state.genes, (g, i) => (
              <Gene value={g} key={i} onChange={this.changeGene(i)} />
            ))}
          </FormControl>
        </Grid>
        <Grid item xs={12} sm={8}>
          <Expression expressionMap={this.state.expressionMap} />
        </Grid>
      </Grid>
    );
  }
}

Show.propTypes = {
  classes: PropTypes.object.isRequired,
  speciesName: PropTypes.string.isRequired,
};

export default withStyles(styles)(Show);
