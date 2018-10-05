import { Button, Card, CardActions, CardContent, CardHeader, Fade } from '@material-ui/core';
import { withStyles } from '@material-ui/core/styles';
import classnames from 'classnames';
import fetch from 'fetch-hoc';
import inflection from 'inflection';
import PropTypes from 'prop-types';
import React from 'react';
import { compose } from 'redux';
import SelectCheck from '../App/SelectCheck';
import WithLoading from '../App/WithLoading';
import { ChromosomeShow } from '../Chromosome';

const _ = require('underscore');


const styles = theme => ({
  root: {
    margin: theme.spacing.unit,
    transition: `all ${theme.transitions.duration.standard}ms ${theme.transitions.easing.easeInOut}`,
  },
  contract: {
    height: 150,
    width: 250,
  },
  expand: {
    height: 800,
    width: 500,
  },
  selected: {
    border: `2px solid ${theme.palette.primary.light}`,
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
      selected: false,
    };
    this.handleToggle = this.handleToggle.bind(this);
    this.handleSelect = this.handleSelect.bind(this);
  }

  handleToggle() {
    this.setState(state => ({ expanded: !state.expanded }));
  }

  handleSelect(selected) {
    this.setState({ selected });
  }

  render() {
    const {
      data, classes, error,
    } = this.props;
    const { expanded, selected } = this.state;
    const { name } = data;
    if (Object.keys(error).length > 0) {
      return (<div>{error.message}</div>);
    }
    const rootClasses = classnames([classes.root, { [classes.selected]: selected, [classes.expand]: expanded, [classes.contract]: !expanded }]);
    return (
      <Card className={rootClasses}>
        <CardHeader
          title={<SelectCheck label={inflection.titleize(name)} onChange={this.handleSelect} />}
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

