import React from 'react';
import Paper from 'material-ui/Paper';
import { withStyles } from 'material-ui/styles';
import IconButton from 'material-ui/IconButton';
import TextField from 'material-ui/TextField';
import AutoRenewIcon from 'material-ui-icons/Autorenew';
import Gene from './Gene';
import chromosomeApi from './api';

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
  render({ classes }) {
    return (
      <div className="flex-container">
        { _.map(this.state.genes, (g, i) => (
          <Gene value={g} onChange={this.changeGene(i)} />
        ))}
        <Gene value={this.state.value} className={classes.textField} onChange={this.handleChange('name')} />
        <IconButton className={classes.avatar} onClick={this.clickRandomChromosome}>
          <AutoRenewIcon />
        </IconButton>
      </div>
    );
  }
}

export default withStyles(styles)(Show);
