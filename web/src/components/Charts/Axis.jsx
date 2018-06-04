import * as d3Axis from 'd3-axis';
import { select as d3Select } from 'd3-selection';
import PropTypes from 'prop-types';
import React from 'react';

class Axis extends React.Component {
  componentDidMount() {
    this.renderAxis();
  }

  componentDidUpdate() {
    this.renderAxis();
  }

  renderAxis() {
    const axisType = `axis${this.props.orient}`;
    const axis = d3Axis[axisType]()
      .scale(this.props.scale)
      .tickSize(-this.props.tickSize)
      .tickPadding([12])
      .ticks([4]);

    d3Select(this.axisElement).call(axis);
  }

  render() {
    const { translate, orient } = this.props;
    return (
      <g
        className={`Axis Axis-${orient}`}
        ref={(el) => { this.axisElement = el; }}
        transform={translate}
      />
    );
  }
}

Axis.propTypes = {
  translate: PropTypes.string.isRequired,
  orient: PropTypes.string.isRequired,
  tickSize: PropTypes.number.isRequired,
  scale: PropTypes.func.isRequired,
};

Axis.defaultProps = {
};

module.exports = Axis;
