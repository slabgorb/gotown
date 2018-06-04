import { scaleBand, scaleLinear } from 'd3-scale';
import PropTypes from 'prop-types';
import React from 'react';
import Axes from './Axes';
import Bars from './Bars';

class Histogram extends React.Component {
  constructor(props) {
    super(props);
    this.xScale = scaleBand();
    this.yScale = scaleLinear();
  }
  render() {
    const { data } = this.props;
    const maxValue = Math.max(...data.map(d => d.value));
    const margins = {
      top: 50,
      right: 20,
      bottom: 100,
      left: 60,
    };
    const svgDimensions = { width: 800, height: 500 };
    const xScale = this.xScale
      .padding(0.5)
      .domain(data.map(d => d.title))
      .range([margins.left, svgDimensions.width - margins.right]);
    const yScale = this.yScale
      .domain([0, maxValue])
      .range([svgDimensions.height - margins.bottom, margins.top]);
    return (
      <svg width={svgDimensions.width} height={svgDimensions.height}>
        <Axes
          scales={{ xScale, yScale }}
          margins={margins}
          svgDimensions={svgDimensions}
        />
        <Bars
          scales={{ xScale, yScale }}
          margins={margins}
          data={data}
          svgDimensions={svgDimensions}
        />
      </svg>
    );
  }
}

Histogram.propTypes = {
  data: PropTypes.array.isRequired,

};
Histogram.defaultProps = {};

module.exports = Histogram;
