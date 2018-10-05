import { Avatar, Button, Card, CardActions, CardContent, CardHeader, Fade } from '@material-ui/core';
import { withStyles } from '@material-ui/core/styles';
import fetch from 'fetch-hoc';
import inflection from 'inflection';
import PropTypes from 'prop-types';
import React from 'react';
import { compose } from 'redux';
import WithLoading from '../App/WithLoading';
import { ChromosomeShow } from '../Chromosome';

const _ = require('underscore');

const trans = theme => `all ${theme.transitions.duration.standard}ms ${theme.transitions.easing.easeInOut}`;

const styles = theme => ({
  contracted: {
    height: 150,
    width: 250,
    transition: trans(theme),
  },
  expanded: {
    height: 800,
    width: 500,
    transition: trans(theme),
  },
  avatar: {
    backgroundColor: theme.palette.primary.main,
  },
});

class SpeciesCard extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      expanded: false,
    };
    this.handleToggle = this.handleToggle.bind(this);
  }

  handleToggle() {
    this.setState(state => ({ expanded: !state.expanded }));
  }

  render() {
    const {
      data, classes, error,
    } = this.props;
    const { expanded } = this.state;
    const { name } = data;
    if (Object.keys(error).length > 0) {
      return (<div>{error.message}</div>);
    }
    return (
      <Card className={expanded ? classes.expanded : classes.contracted}>
        <CardHeader

          avatar={(<Avatar area-label="Species" className={classes.avatar}>S</Avatar>)}
          title={inflection.titleize(name)}
        />
        <CardActions>
          <Button onClick={this.handleToggle}>Chromosome</Button>
        </CardActions>
        <Fade in={expanded}>
          <CardContent>
            <ChromosomeShow speciesName={name} speciesID={data.id} traits={data.expression.traits} />
          </CardContent>
        </Fade>
      </Card>
    );
  }
}

SpeciesCard.propTypes = {
  data: PropTypes.object,
  classes: PropTypes.object.isRequired,
  loading: PropTypes.bool,
  error: PropTypes.string,
};

SpeciesCard.defaultProps = {
  error: '',
  data: {},
  loading: true,
};

module.exports = compose(
  fetch(({ id }) => `/api/species/${id}`),
  WithLoading,
  withStyles(styles),
)(SpeciesCard);

