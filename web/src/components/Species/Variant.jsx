import React from 'react';
import PropTypes from 'prop-types';

const Variant = ({ name, match }) =>
  (
    <div className="key-value">
      <div>{name}</div>
      <div>{match}</div>
    </div>
  );

Variant.propTypes = {
  name: PropTypes.string.isRequired,
  match: PropTypes.string.isRequired,
};

module.exports = Variant;
