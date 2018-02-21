import React from 'react';


const hexes = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'];
const blankLines = h =>
  (
    <div key={h}>
      <span>{h}</span>
      { hexes.map(hx => (<span key={hx}>&nbsp;</span>))}
    </div>
  );

const line = hexes.map(h => (<span>{h}</span>));


const GeneticsMap = () =>
  (
    <div>
      <span>&nbsp;</span>{line}
      { hexes.map(h => blankLines(h)) }
    </div>
  );

module.exports = GeneticsMap;
