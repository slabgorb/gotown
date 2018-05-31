import Card from '@material-ui/core/Card';
import PropTypes from 'prop-types';
import React from 'react';

const hexes = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'];
const lines = h => (
  <div key={h} className="tiny-box-row">
    <div className="tiny-box" >{h}</div>
    { hexes.map(hx => (<div className="tiny-box" key={hx}>&nbsp;</div>))}
  </div>
);


const line = hexes.map(h => (<div key={h} className="tiny-box">{h}</div>));


const GeneticsMap = ({ traits }) =>
  (
    <Card>
      <div className="tiny-box-container">
        <div className="tiny-box-row"><div className="tiny-box" >&nbsp;</div>{line}</div>
        { hexes.map(h => lines(h, traits)) }
      </div>
    </Card>
  );

GeneticsMap.propTypes = {
  traits: PropTypes.array.isRequired,
};

module.exports = GeneticsMap;
