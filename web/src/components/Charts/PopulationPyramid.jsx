import { scaleBand, scaleLinear } from 'd3-scale';
import PropTypes from 'prop-types';
import React from 'react';
import Axes from './Axes';
import Bars from './Bars';

const genderFilter = data => gender => data.filter(d => d.sex === gender);
const mapToSegments = (data, segments) => segments.reduce()

class PopulationPyramid extends React.Component {
  constructor(props) {
    super(props);
    this.xScaleFemale = scaleLinear();
    this.xScaleMale = scaleLinear();
    this.yScale = scaleBand();
  }

  render() {
    const {
      data,
      width,
      height,
      segments,
    } = this.props;
    const margins = {
      top: 50,
      right: 20,
      bottom: 100,
      left: 60,
    };
    const yScale = this.yScale
      .domain(segments.map(d => d.title))
      .range([height - margins.bottom, margins.top]);
    const maxValue = gender => Math.max(...genderFilter(data)(gender).map(d => d.age));
    const xScaleFemale = this.xScaleFemale
      .domain([0, maxValue('female')])
      .range([margins.left, width - margins.right]);
    const xScaleMale = this.xScaleMale
      .domain([0, maxValue('male')])
      .range([margins.left, width - margins.right]);
    return (
      <svg width={width} height={height}>
        <Axes
          scales={{ xScale: xScaleFemale, yScale }}
          margins={margins}
          svgDimensions={{ width, height }}
        />
        <Axes
          scales={{ xScale: xScaleMale, yScale }}
          margins={margins}
          svgDimensions={{ width, height }}
        />
        <Bars
          scales={{ xScale: xScaleFemale, yScale }}
          margins={margins}
          data={genderFilter(data)('female')}
          svgDimensions={{ width, height }}
        />
        <Bars
          scales={{ xScale: xScaleMale, yScale }}
          margins={margins}
          data={genderFilter(data)('male')}
          svgDimensions={{ width, height }}
        />
      </svg>

    )
  }

}

PopulationPyramid.propTypes = {
  data: PropTypes.array.isRequired,
  segments: PropTypes.array.isRequired,
  width: PropTypes.number,
  height: PropTypes.number,
}

PopulationPyramid.defaultProps = {
  width: 300,
  height: 450,
}

module.exports = PopulationPyramid;