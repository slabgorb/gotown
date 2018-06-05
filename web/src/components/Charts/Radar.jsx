import { scaleBand, scaleLinear } from 'd3-scale';
import PropTypes from 'prop-types';
import React from 'react';

const radians = 2 * Math.PI;

class Radar extends React.Component {
  constructor(props) {
    super(props);
    this.xScale = scaleBand();
    this.yScale = scaleLinear();
    this.axis = this.axis.bind(this);
    this.getPosition = this.getPosition.bind(this);
  }
  
  getPosition(i, rnge, func) {
    const total = this.props.data[0].axes.length;
    const { factor } = this.props;
    const pos =  rnge * (1 - (factor * func((i * radians) / total)));
    return pos
  }
  getHorizontalPosition(i, rnge) {
    return this.getPosition(i, rnge, Math.sin);
  }
  getVerticalPosition(i, rnge) {
    return this.getPosition(i, rnge, Math.cos);
  }

  axis() {
    const {
      data,
      w,
      h,
      factor,
    } = this.props;
    const outerRadius = Math.min(w / 2, h / 2);
    const axes = data[0].axes.map(i => ({ name: i.axis, xOffset: (i.xOffset) ? i.xOffset : 0, yOffset: (i.yOffset) ? i.yOffset : 0 }));
    const axisLine = (x1, y1, x2, y2) => (
      <line
        key={`${x1}${x2}${y1}${y2}`}
        x1={x1}
        y1={y1}
        x2={x2}
        y2={y2}
        fill="black"
        stroke="black"
      />
    );
    const axisLabel = (x1, y1, text) => (
      <text
        key={text}
        x={x1}
        y={y1}
        fill="black"
        stroke="black"
      >
        {text}
      </text>
    );
    const children = axes.map((a, i) =>
      axisLine(
        w / 2,
        h / 2,
        ((w / 2) - outerRadius) + this.getHorizontalPosition(i, outerRadius, factor),
        ((h / 2) - outerRadius) + this.getVerticalPosition(i, outerRadius, factor),
      ));
    const labels = axes.map((a, i) =>
      axisLabel(
        ((w / 2) - outerRadius) + this.getHorizontalPosition(i, outerRadius, factor),
        ((h / 2) - outerRadius) + this.getVerticalPosition(i, outerRadius, factor),
        a.axis,
      ));
    return (<g className="RadarAxis">{children}{labels}</g>);
  }

  render() {
    const {
      data, w, h, factor, levels,
    } = this.props;
    const maxValue = Math.max(...data.map(d => d.value));
    const radius = factor * Math.min(w / 2, h / 2);

    const margins = {
      top: 50,
      right: 20,
      bottom: 100,
      left: 60,
    };

    //let levelFactors = range(0, levels).map(level => radius * ((level + 1) / levels));



    return (
      <svg width={w} height={h}>
        {this.axis()}
      </svg>
    );
  }
}

Radar.propTypes = {
  data: PropTypes.array.isRequired,
  w: PropTypes.number,
  h: PropTypes.number,
  factor: PropTypes.number,
  levels: PropTypes.number,
};

Radar.defaultProps = {
  w: 500,
  h: 500,
  factor: 0.95,
  levels: 3,
}

module.exports = Radar;