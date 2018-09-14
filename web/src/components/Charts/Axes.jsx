import PropTypes from 'prop-types';
import React from 'react';
import Axis from './Axis';

const Axes = ({
  scales,
  margins,
  svgDimensions,
  ...props
}) => {
  const { height, width } = svgDimensions;

  const xProps = {
    orient: 'Bottom',
    scale: scales.xScale,
    translate: `translate(0, ${height - margins.bottom})`,
    tickSize: height - margins.top - margins.bottom,
  };

  const yProps = {
    orient: 'Left',
    scale: scales.yScale,
    translate: `translate(${margins.left}, 0)`,
    tickSize: width - margins.left - margins.right,
  };

  return (
    <g {...props} >
      <Axis {...xProps} />
      <Axis {...yProps} />
    </g>
  );
};

Axes.propTypes = {
  scales: PropTypes.object.isRequired,
  margins: PropTypes.object.isRequired,
  svgDimensions: PropTypes.object.isRequired,
};

module.exports = Axes;
