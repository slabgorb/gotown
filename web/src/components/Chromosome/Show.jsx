import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import FormLabel from '@material-ui/core/FormLabel';
import Grid from '@material-ui/core/Grid';
import IconButton from '@material-ui/core/IconButton';
import { withStyles } from '@material-ui/core/styles';
import AutoRenewIcon from '@material-ui/icons/Autorenew';
import PropTypes from 'prop-types';
import React from 'react';
import speciesApi from '../Species/api';
import chromosomeApi from './api';
import Expression from './Expression';
import Gene from './Gene';

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
      <Grid
        container
        spacing={24}
        className={classes.root}
        xs={9}
        sm={9}
        lg={9}
      >
        <Grid container item xs={12}>
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
              {this.state.genes.length == 0 ? "" :
                this.state.genes.map((g, i) => (
                  <Button
                    onClick={this.clickGeneInMap(i)}
                    className="gene"
                    style={{
                      borderBottomColor: `#${g}`,
                      borderBottomWidth: '2px',
                      borderBottomStyle: 'solid',
                    }}
                    key={g}
                  >
                    {g}
                  </Button>))
              }
            </div>
          </Grid>
        </Grid>
        <Grid container xs={12}>
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
