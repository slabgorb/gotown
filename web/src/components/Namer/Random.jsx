import AutoRenewIcon from 'material-ui-icons/Autorenew';
import Grid from 'material-ui/Grid';
import IconButton from 'material-ui/IconButton';
import Typography from 'material-ui/Typography';
import { withStyles } from 'material-ui/styles';
import PropTypes from 'prop-types';
import React from 'react';
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
      name: props.name,
      id: props.id,
    };
    this.get = this.get.bind(this);
  }

  componentWillMount() {
    this.get();
  }

  get() {
    api.random(this.state.id)
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
  id: PropTypes.number.isRequired,
  name: PropTypes.string,
  classes: PropTypes.object.isRequired,
};

Random.defaultProps = {
  name: '',
};

module.exports = withStyles(style)(Random);
