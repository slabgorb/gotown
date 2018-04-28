import React from 'react';
import PropTypes from 'prop-types';
import Typography from 'material-ui/Typography';
import { withStyles } from 'material-ui/styles';
import AutoRenewIcon from 'material-ui-icons/Autorenew';
import IconButton from 'material-ui/IconButton';
import Grid from 'material-ui/Grid';
import api from './api';

const style = () => ({
  root: {},
  avatar: {
    width: '100%',
  },
  box: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
});

class Random extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      namer: props.namer,
      name: props.name,
    };
    this.get = this.get.bind(this);
  }

  componentWillMount() {
    this.get();
  }

  get() {
    api.random(this.state.namer)
      .then(name => this.setState({ name }));
  }

  render() {
    const { name } = this.state;
    const { classes } = this.props;
    return (
      <Grid container className={classes.box}>
        <Typography
          variant="display2"
          align="center"
          color="secondary"
        >
          {name}
        </Typography>
        <IconButton
          className={classes.avatar}
          onClick={this.get}
        >
          <AutoRenewIcon />
        </IconButton>
      </Grid>
    );
  }
}

Random.propTypes = {
  namer: PropTypes.string.isRequired,
  name: PropTypes.string,
  classes: PropTypes.object.isRequired,
};

Random.defaultProps = {
  name: '',
};

module.exports = withStyles(style)(Random);
