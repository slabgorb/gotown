import PropTypes from 'prop-types';
import React from 'react';

const _ = require('underscore');

const radians = 2 * Math.PI;

class Radar extends React.Component {
  constructor(props) {
    super(props);


    this.axis = this.axis.bind(this);
    this.getPosition = this.getPosition.bind(this);
  }
  
  getPosition(i, rnge, fact, func) {
    const { data } = this.props;
    const total = data[0].axes.length;
    const pos = rnge * (1 - (fact * func((i * radians) / total)));
    return pos;
  }
  getHorizontalPosition(i, rnge, fact) {
    return this.getPosition(i, rnge, fact, Math.sin);
  }
  getVerticalPosition(i, rnge, fact) {
    return this.getPosition(i, rnge, fact, Math.cos);
  }

  maxValue() {
    const { data } = this.props;
    return Math.max(..._.flatten(data.map(d => _.map(d.axes, a => a.value))))
  }

  outerRadius() {
    const { w, h } = this.props;
    return Math.min(w / 2, h / 2);
  }

  circles() {
    const {
      w,
      h,
      factor,
    } = this.props;
    const cx = w / 2;
    const cy = h / 2;
    const outerRadius = this.outerRadius();
    const circleCount = 2;
    const circs = [];
    for (let i = circleCount; i > 0; i -= 1) {
      const radius = (outerRadius / i) * factor;
      circs.push((
        <circle
          key={i}
          cx={cx}
          cy={cy}
          r={radius}
          stroke="#E0E0E0"
          fill="transparent"
        />
      ));
    } 
    return (<g>{circs}</g>)
  }

  chart() {
    const {
      data,
      factor,
    } = this.props;
    const dots = _.flatten(_.map(data, v =>
      _.map(v.axes, (a, i) => {
        const fact = (a.value / this.maxValue()) * factor;
        return (
          <circle
            key={i}
            cx={this.getHorizontalPosition(i, this.outerRadius(), fact)}
            cy={this.getVerticalPosition(i, this.outerRadius(), fact)}
            r={5}
          />
        );
      })));
    const polyPoints = _.flatten(_.map(data, v =>
      _.map(v.axes, (a, i) => {
        const fact = (a.value / this.maxValue()) * factor;
        return [this.getHorizontalPosition(i, this.outerRadius(), fact), this.getVerticalPosition(i, this.outerRadius(), fact)].join(',');
      }))).join(' ');
    return (
      <g className="RadarChart">
        <polygon points={polyPoints} />
        {dots}
      </g>
    );
  }

  axis() {
    const {
      data,
      w,
      h,
      factor,
    } = this.props;
    const outerRadius = this.outerRadius();
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
        a.name,
      ));
    return (<g className="RadarAxis">{children}{labels}</g>);
  }

  render() {
    const {
      w, h,
    } = this.props;
    return (
      <svg width={w} height={h}>
        {this.axis()}
        {this.circles()}
        {this.chart()}
      </svg>
    );
  }
}

Radar.propTypes = {
  data: PropTypes.array.isRequired,
  w: PropTypes.number,
  h: PropTypes.number,
  factor: PropTypes.number,
};

Radar.defaultProps = {
  w: 500,
  h: 500,
  factor: 0.85,
}

module.exports = Radar;