import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';
import IconButton from 'material-ui/IconButton';
import AutoRenewIcon from 'material-ui-icons/Autorenew';
import Gene from './Gene';
import chromosomeApi from './api';

const _ = require('underscore');

const styles = () => ({
  textField: {},
  avatar: {},
});

class Show extends React.Component {
 
  constructor(props) {
    super(props);
    this.state = {
      genes: [],
    };
  }

  changeGene(index) {
    return (value) => {
      this.setState(prevState => ({
        genes: [...prevState.genes.slice(0, index - 1), value, ...prevState.genes.slice(index + 1)],
      }));
    };
  }

  clickRandomChromosome() {
    chromosomeApi.random().then(({ genes }) => {
      this.setState({
        genes,
      });
    });
  }
  render() {
    const { classes } = this.props;
    return (
      <div className="flex-container">
        { _.map(this.state.genes, (g, i) => (
          <Gene value={g} onChange={this.changeGene(i)} />
        ))}
        <IconButton className={classes.avatar} onClick={this.clickRandomChromosome}>
          <AutoRenewIcon />
        </IconButton>
      </div>
    );
  }
}

Show.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(Show);
