import { Avatar, Card, CardContent, CardHeader, Fade } from '@material-ui/core';
import { withStyles } from '@material-ui/core/styles';
import fetch from 'fetch-hoc';
import inflection from 'inflection';
import PropTypes from 'prop-types';
import React from 'react';
import { compose } from 'redux';
import WithLoading from '../App/WithLoading';
import { ChromosomeShow } from '../Chromosome';

const _ = require('underscore');

const trans = theme => 'all 450ms cubic-bezier(0.23, 1, 0.32, 1)';

const styles = theme => ({
  contracted: {
    height: 100,
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
          onClick={this.handleToggle}
          avatar={(<Avatar area-label="Species" className={classes.avatar}>S</Avatar>)}
          title={inflection.titleize(name)}
        />
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

