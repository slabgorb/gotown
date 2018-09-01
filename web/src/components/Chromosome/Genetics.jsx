import React from 'react';
import PropTypes from 'prop-types';
import Trait from './Trait';

const Genetics = ({ traits }) =>
  (
    <div>
      { traits.map(trait =>
        (<Trait key={trait.name} {...trait} />))
      }
    </div>
  );

Genetics.propTypes = {
  traits: PropTypes.array.isRequired,
};

module.exports = Genetics;
