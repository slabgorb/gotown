import { withStyles } from '@material-ui/core/styles';
import PropTypes from 'prop-types';
import React from 'react';

const styles = () => ({
  display: {
    width: '100%',
  },
  displayItem: {
    display: 'inline',
  }

});

const colors = ['#757ce8','#3f50b5','#002884']

class BlendedChooser extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      percents: [0],
    };
    this.getStyle = this.getStyle.bind(this);
  }

  getStyle(i) {
    return ({
      color: colors[i],
      width: `${this.state.percents[i]}%`,
    });
  }

  render() {
    const { classes, choices } = this.props;
    return (
      <React.Fragment>
        {choices}
        <div className={classes.display}>
          { choices.map((_,i) => <div className={classes.displayItem} style={this.getStyle(i)} />) }
        </div>
      </React.Fragment>
    );
  }
}

BlendedChooser.propTypes = {
  classes: PropTypes.object.isRequired,
  choices: PropTypes.array.isRequired,
};

export default withStyles(styles)(BlendedChooser);

