import Grid from '@material-ui/core/Grid';
import { withStyles } from '@material-ui/core/styles';
import PropTypes from 'prop-types';
import React from 'react';
import { SpeciesCard } from '../Species';
import Info from './Info';
import PageTitle from './PageTitle';

const _ = require('underscore');

const style = () => ({
  info: {
    marginTop: 48,
  },
  root: {},
});

const Home = ({ classes }) => {
  const infos = {
    'Ever needed to generate a quick village in a role playing game?': {
      paras: [
        'Your party has to run away to somewhere. Or you need a place to hook the next adventure in. Or you are looking for a town setting for an adventure.',
        'We have you covered.',
      ],
    },
    'Cultures and Subcultures': {
      paras: [
        'Cultures define the way towns and people are named, as well as their marital customs.',
      ],
    },
    'Species and Subraces': { 
      paras: [
        'Species and subraces can be defined with individualized traits.'
      ],
      children: (
        <Grid container>
          <SpeciesCard id="d54f106f-d65e-4ead-93f9-3dd775be95ce" />
        </Grid>

      ),
    },
    'Genetics and Stereotypes': { 
      paras: [
        'People (and whatever you want) are generated using a sort of genetic code, which works with the traits of the species/subrace to generate a randomized but consistent set of traits.',
        'You can define a town using a chosen set of base genetic codes, which are mutated slightly.',
        'This allows towns to have people which share similar traits, and gives a basis for compariason between towns and individuals',
      ],
    },
    'Heraldry Icons': {
      paras: [
        'To give a sort of accent to the settlement, you are provided a randomized bit of heraldry.'
      ],
      children: (
        <Grid container>
          <Grid item xs={3}>
            <img src="/api/random/heraldry?a=1" alt="random" width={64} height={64} />
          </Grid>
          <Grid item xs={3}>
            <img src="/api/random/heraldry?a=2" alt="random" width={64} height={64} />
          </Grid>
          <Grid item xs={3}>
            <img src="/api/random/heraldry?a=3" alt="random" width={64} height={64} />
          </Grid>
          <Grid item xs={3}>
            <img src="/api/random/heraldry?a=4" alt="random" width={64} height={64} />
          </Grid>
        </Grid>
      ),
    },
  };
  return (
    <div>
      <PageTitle title="Home" titleize classes={classes} />
      <Grid container className={classes.info}>
        {_.map(infos, (info, header) => (<Info key={header} header={header} paras={info.paras}>{info.children}</Info>))}
      </Grid>
    </div>
  );
};

Home.propTypes = {
  classes: PropTypes.object.isRequired,
};

module.exports = withStyles(style)(Home);
