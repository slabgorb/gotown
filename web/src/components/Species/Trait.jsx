import React from 'react';
import Card from 'material-ui/Card';
import PropTypes from 'prop-types';

const Trait = ({ name, variants }) =>
  (
    <div>
      <Card>
        <h3>{name}</h3>
        {variants.map((variant) => (<Variant name={variant.name} match={variant.match} key={variant.name}/>))}
      </Card>
    </div>
  );


Trait.propTypes = {
  name: PropTypes.string.isRequired,
  variants: PropTypes.object.isRequired,
};

module.exports = Trait;
