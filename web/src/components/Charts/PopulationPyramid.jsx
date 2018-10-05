import { scaleBand, scaleLinear } from 'd3-scale';
import PropTypes from 'prop-types';
import React from 'react';
import Axes from './Axes';
import Bars from './Bars';

const genderFilter = data => gender => data.filter(d => d.gender === gender);
const mapToSegments = (data, segments) => {
  const endData = [];
  let minAge = 0;
  const filterFunc = (min, max) => d => d.age >= min && d.age <= max;
  for (let i = 0; i < segments.length; i += 1) {
    endData[i] = {
      title: `${minAge} to ${segments[i].maxAge}`,
      value: data.filter(filterFunc(minAge, segments[i].maxAge)).length,
    };
    minAge = segments[i].maxAge;
  }
  return endData;
};

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
    const maleData = mapToSegments(genderFilter(data)('male'), segments);
    const femaleData = mapToSegments(genderFilter(data)('female'), segments);
    console.log(maleData)
    console.log(femaleData)
    const axes = (
      <g>
        <Axes
          scales={{ xScale: xScaleFemale, yScale }}
          margins={margins}
          svgDimensions={{ width: width / 2, height }}
          transform={`translate(${width / 2}, 0)`}
        />
        <Axes
          scales={{ xScale: xScaleMale, yScale }}
          margins={margins}
          svgDimensions={{ width: width / 2, height }}
        />
      </g>
    );
    return (
      <svg width={width} height={height}>
        {axes}
        <Bars
          scales={{ xScale: xScaleFemale, yScale }}
          margins={margins}
          data={femaleData}
          svgDimensions={{ width: width / 2, height }}
          transform={`translate(${width / 2}, 0)`}
        />
        <Bars
          scales={{ xScale: xScaleMale, yScale }}
          margins={margins}
          data={maleData}
          svgDimensions={{ width: width / 2, height }}
        />
      </svg>

    );
  }
}

PopulationPyramid.propTypes = {
  data: PropTypes.array.isRequired,
  segments: PropTypes.array.isRequired,
  width: PropTypes.number,
  height: PropTypes.number,
};

PopulationPyramid.defaultProps = {
  width: 300,
  height: 450,
};

module.exports = PopulationPyramid;
