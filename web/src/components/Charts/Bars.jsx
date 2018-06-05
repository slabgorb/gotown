import PropTypes from 'prop-types';
import React from 'react';

const Bars = ({
  scales,
  margins,
  data,
  svgDimensions,
}) => {
  const { xScale, yScale } = scales;
  const { height } = svgDimensions;

  const bars = (
    data.map(datum =>
      (<rect
        key={datum.title}
        x={xScale(datum.title)}
        y={yScale(datum.value)}
        height={height - margins.bottom - scales.yScale(datum.value)}
        width={xScale.bandwidth()}
        fill="blue"
      />))
  );

  return (
    <g>{bars}</g>
  );
};


Bars.propTypes = {
  data: PropTypes.array.isRequired,
  scales: PropTypes.object.isRequired,
  margins: PropTypes.object.isRequired,
  svgDimensions: PropTypes.object.isRequired,
};

module.exports = Bars;
