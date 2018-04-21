import React from 'react';
import { withStyles } from 'material-ui/styles';
import PropTypes from 'prop-types';
import PageTitle from './PageTitle';

const style = () => ({

});

const Home = ({ classes }) => (
  <PageTitle title="Home" titleize classes={classes} />
);


Home.propTypes = {
  classes: PropTypes.object.isRequired,
};

module.exports = withStyles(style)(Home);
