import { Checkbox, FormControlLabel } from '@material-ui/core';
import { withStyles } from '@material-ui/core/styles';
import PropTypes from 'prop-types';
import React from 'react';


const styles = () => ({});

class SelectCheck extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      checked: props.defaultChecked,
    };
    this.handleChange = this.handleChange.bind(this);
  }

  handleChange(e) {
    const { onChange } = this.props;
    this.setState({ checked: e.target.checked });
    onChange(e.target.checked);
  }

  render() {
    const { checked } = this.state;
    const { classes, label } = this.props;
    return (
      <FormControlLabel
        control={
          <Checkbox checked={checked} classNames={classes.check} onClick={this.handleChange} />
        }
        label={label}
      />
    );
  }
}


SelectCheck.propTypes = {
  defaultChecked: PropTypes.bool,
  onChange: PropTypes.func.isRequired,
  classes: PropTypes.object.isRequired,
  label: PropTypes.string,
};

SelectCheck.defaultProps = {
  defaultChecked: false,
  label: '',
};


export default withStyles(styles)(SelectCheck);
