import React from 'react';
import PropTypes from 'prop-types';
import Trait from './Trait';

const Genetics = ({ traits }) =>
  (
    <div>
      { traits.map(trait => (<Trait name={trait.name} key={trait.name} variants={trait.variants}/>)) }
    </div>
  );

Genetics.propTypes = {
  traits: PropTypes.object.isRequired,
};

module.exports = Genetics;
