import React from 'react';
import { withStyles } from 'material-ui/styles';
import PropTypes from 'prop-types';
import Typography from 'material-ui/Typography';
import Drawer from 'material-ui/Drawer';
import AppBar from 'material-ui/AppBar';
import Toolbar from 'material-ui/Toolbar';
import IconButton from 'material-ui/IconButton';
import Button from 'material-ui/Button';
import MenuIcon from 'material-ui-icons/Menu';
import CloseIcon from 'material-ui-icons/Close';
import Grid from 'material-ui/Grid';
import Dialog, { DialogTitle } from 'material-ui/Dialog';
import { withRouter } from 'react-router-dom';
import inflection from 'inflection';
import axios from 'axios';
import { SpeciesList } from '../Species';
import { CulturesList } from '../Culture';
import { WordsList } from '../Words';
import { NamersList } from '../Namer';
import { AreaList } from '../Area';


const seed = () => axios.put('/api/seed');

const style = theme => ({
  flex: { flex: '1 0' },
  pageRoot: { marginBottom: 64 },
  appBar: {},
  root: {
    transition: '0.3s all',
    backgroundColor: theme.palette.primary.main,
    color: 'white',
  },
  subtitle: {
    backgroundColor: theme.palette.primary.main,
    color: 'black',
    marginTop: '-1ex',
    paddingRight: '4rem',
  },
});

class PageTitle extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      drawerOpen: false,
      dialogOpen: false,
    };
    this.handleDrawer = this.handleDrawer.bind(this);
    this.handleMenuItem = this.handleMenuItem.bind(this);
    this.handleDialog = this.handleDialog.bind(this);
  }

  handleDialog(dialogOpen) {
    return () => {
      this.setState({ dialogOpen });
    };
  }

  handleDrawer(drawerOpen) {
    return () => {
      this.setState({ drawerOpen });
    };
  }

  handleMenuItem(value) {
    const { history } = this.props;
    return () => {
      history.push(`/${value}`);
      this.handleDrawer(false)();
    };
  }

  render() {
    const {
      classes,
      title,
      titleize,
      capitalize,
      subtitle,
    } = this.props;
    const { drawerOpen, dialogOpen } = this.state;
    let endTitle = title;
    if (capitalize) { endTitle = inflection.capitalize(title); }
    if (titleize) { endTitle = inflection.titleize(title); }
    return (
      <div className={classes.pageRoot}>
        <Dialog open={dialogOpen}>
          <DialogTitle>Are you sure?</DialogTitle>
          <Typography>
            This will destroy the current database and recreate
            it will seeded data.
          </Typography>
          <Toolbar>
            <Button onClick={() => seed().then(this.handleDialog(false))}>Yes</Button>
            <Button onClick={this.handleDialog(false)}>Cancel</Button>
          </Toolbar>
        </Dialog>
        <Drawer
          open={drawerOpen}
          anchor="right"
          onClose={this.handleDrawer(false)}
        >
          <div>
            <AppBar position="static" className={classes.appBar}>
              <Toolbar>
                <IconButton color="default" onClick={this.handleDrawer(false)}>
                  <CloseIcon className={classes.button} />
                </IconButton>
                <Button className={classes.button} onClick={this.handleDialog(true)} >Seed</Button>
              </Toolbar>
            </AppBar>
            <SpeciesList handleClick={v => this.handleMenuItem(v)} />
            <CulturesList handleClick={v => this.handleMenuItem(v)} />
            <WordsList handleClick={v => this.handleMenuItem(v)} />
            <NamersList handleClick={v => this.handleMenuItem(v)} />
            <AreaList handleClick={v => this.handleMenuItem(v)} />
          </div>
        </Drawer>
        <AppBar position="fixed" className={classes.appBar}>
          <Toolbar>
            <Grid container alignContent="space-between" alignItems="center" spacing={24}>
              <Grid xs={3} item>
                <Typography variant="title" color="inherit" className={classes.flex}>
                  <a className="logo" href="/">Gotown</a>
                </Typography>
                { subtitle !== '' && (
                  <Typography variant="subheading" align="right" classes={{ root: classes.subtitle }}>
                    {subtitle}
                  </Typography>
                )}
              </Grid>
              <Grid xs={8} item>
                <Typography variant="display1" align="right" classes={{ root: classes.root }}>
                  {endTitle}
                </Typography>
              </Grid>
            </Grid>
            <Grid xs={1} item>
              <IconButton onClick={this.handleDrawer(true)}>
                <MenuIcon />
              </IconButton>
            </Grid>
          </Toolbar>
        </AppBar>
      </div>
    );
  }
}

PageTitle.propTypes = {
  history: PropTypes.object.isRequired,
  classes: PropTypes.object.isRequired,
  title: PropTypes.string.isRequired,
  titleize: PropTypes.bool,
  capitalize: PropTypes.bool,
  subtitle: PropTypes.string,
};

PageTitle.defaultProps = {
  titleize: false,
  capitalize: false,
  subtitle: '',
};

module.exports = withRouter(withStyles(style)(PageTitle));
