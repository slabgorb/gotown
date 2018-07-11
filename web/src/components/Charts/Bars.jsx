import PropTypes from 'prop-types';
import React from 'react';

const Bars = ({
  scales,
  margins,
  data,
  svgDimensions,
  orientation,
}) => {
  const { xScale, yScale } = scales;
  const { height } = svgDimensions;

  const bars = (
    data.map(datum =>
      (<rect
        key={datum.title}
        x={orientation === 'horizontal' ? xScale(datum.title) : yScale(datum.value)}
        y={orientation === 'horizontal' ? yScale(datum.value) : xScale(datum.title)}
        height={height - margins.bottom - (orientation === 'horizontal' ? scales.yScale(datum.value) : scales.yScale(datum.title))}
        width={orientation === 'horizontal' ? xScale.bandwidth() : yScale.bandwidth()}
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
  orientation: PropTypes.oneOf(['vertical', 'horizontal']),

};

Bars.defaultProps = {
  orientation: 'vertical',
}

module.exports = Bars;
