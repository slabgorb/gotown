import { withStyles } from '@material-ui/core/styles';
import PropTypes from 'prop-types';
import React from 'react';
import Info from './Info';
import PageTitle from './PageTitle';

const _ = require('underscore');

const style = () => ({

});

const Home = ({ classes }) => {
  const infos = {
    'Ever needed to generate a quick village in a role playing game?': [
      'Your party has to run away to somewhere. Or you need a place to hook the next adventure in. Or you are looking for a town setting for an adventure.',
      'We have you covered.',
    ],
    'Cultures and Subcultures': [
      'Cultures define the way towns and people are named, as well as their marital customs.',
    ],
    'Species/Subraces': [
      'Species and subraces can be defined with individualized traits.'
    ],
    'Genetics and Stereotypes': [
      'People (and whatever you want) are generated using a sort of genetic code, which works with the traits of the species/subrace to generate a randomized but consistent set of traits.',
      'You can define a town using a chosen set of base genetic codes, which are mutated slightly.',
      'This allows towns to have people which share similar traits, and gives a basis for compariason between towns and individuals',
    ],
    'Heraldry Icons': [
      'To give a sort of accent to the settlement, you are provided a randomized bit of heraldry.'
    ]};
  return (
    <div>
      <PageTitle title="Home" titleize classes={classes} />
      {_.map(infos, (paras, header) => (<Info header={header} paras={paras} />))}
    </div>
  );
};

Home.propTypes = {
  classes: PropTypes.object.isRequired,
};

module.exports = withStyles(style)(Home);
