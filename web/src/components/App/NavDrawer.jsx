import AppBar from '@material-ui/core/AppBar';
import Button from '@material-ui/core/Button';
import Drawer from '@material-ui/core/Drawer';
import IconButton from '@material-ui/core/IconButton';
import Toolbar from '@material-ui/core/Toolbar';
import { withStyles } from '@material-ui/core/styles';
import CloseIcon from '@material-ui/icons/Close';
import PropTypes from 'prop-types';
import React from 'react';
import { withRouter } from 'react-router-dom';
import { AreaList } from '../Area';
import { CulturesList } from '../Culture';
import { NamersList } from '../Namer';
import { SpeciesList } from '../Species';
import { WordsList } from '../Words';

const style = theme => ({
  root: {
    transition: '0.3s all',
    backgroundColor: theme.palette.primary.main,
  },
});
class NavDrawer extends React.Component {
  constructor(props) {
    super(props);
    this.handleMenuItem = this.handleMenuItem.bind(this);
  }
  handleMenuItem(value) {
    const { history } = this.props;
    return () => {
      history.push(`/${value}`);
      this.props.handleClose();
    };
  }
  render() {
    const {
      classes,
      open,
      onClose,
      handleDialog,
      handleClose,
    } = this.props;
    return (
      <Drawer
        className={classes.root}
        open={open}
        anchor="right"
        onClose={onClose}
      >
        <div>
          <AppBar position="static" className={classes.appBar}>
            <Toolbar>
              <IconButton color="default" onClick={handleClose}>
                <CloseIcon className={classes.button} />
              </IconButton>
              <Button className={classes.button} onClick={handleDialog} >Seed</Button>
              <Button className={classes.button} onClick={this.handleMenuItem('towns/create')} >New Town</Button>
            </Toolbar>
          </AppBar>
          <SpeciesList handleClick={v => this.handleMenuItem(v)} />
          <CulturesList handleClick={v => this.handleMenuItem(v)} />
          <WordsList handleClick={v => this.handleMenuItem(v)} />
          <NamersList handleClick={v => this.handleMenuItem(v)} />
          <AreaList handleClick={v => this.handleMenuItem(v)} />
        </div>
      </Drawer>
    );
  }
}

NavDrawer.propTypes = {
  history: PropTypes.object.isRequired,
  classes: PropTypes.object.isRequired,
  open: PropTypes.bool,
  onClose: PropTypes.func.isRequired,
  handleDialog: PropTypes.func.isRequired,
  handleClose: PropTypes.func.isRequired,
};

NavDrawer.defaultProps = {
  open: false,
};
module.exports = withRouter(withStyles(style)(NavDrawer));
