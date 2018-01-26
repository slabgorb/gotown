import React from 'react';
import PropTypes from 'prop-types';

const GenderName = ({ gender, patterns, givenNames }) =>
  (
    <div className="gender-name">
      <h3>{gender}</h3>
      <p>{patterns.join(', ')}</p>
      <p>{givenNames.join(', ')}</p>
    </div>
  );


GenderName.propTypes = {
  gender: PropTypes.string.isRequired,
  patterns: PropTypes.object.isRequired,
  givenNames: PropTypes.array.isRequired,
};


module.exports = GenderName
