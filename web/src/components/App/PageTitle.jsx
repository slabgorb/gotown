import AppBar from '@material-ui/core/AppBar';
import Grid from '@material-ui/core/Grid';
import IconButton from '@material-ui/core/IconButton';
import { withStyles } from '@material-ui/core/styles';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import MenuIcon from '@material-ui/icons/Menu';
import axios from 'axios';
import inflection from 'inflection';
import PropTypes from 'prop-types';
import React from 'react';
import { withRouter } from 'react-router-dom';
import NavDrawer from './NavDrawer';
import SimpleDialog from './SimpleDialog';

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

  render() {
    const {
      classes,
      title,
      titleize,
      capitalize,
      subtitle,
      icon,
    } = this.props;
    const { drawerOpen, dialogOpen } = this.state;
    let endTitle = title;
    if (capitalize) { endTitle = inflection.capitalize(title); }
    if (titleize) { endTitle = inflection.titleize(title); }
    return (
      <div className={classes.pageRoot}>
        <SimpleDialog
          open={dialogOpen}
          title="Are you sure?"
          message="This will destroy the current database and recreate it with seeded data."
          onYes={() => seed().then(this.handleDialog(false))}
          onCancel={this.handleDialog(false)}
        />
        <NavDrawer
          open={drawerOpen}
          onClose={this.handleDrawer(false)}
          handleDialog={this.handleDialog(true)}
          handleClose={this.handleDrawer(false)}
        />
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
              {icon && (<Grid xs={1} item>{icon}</Grid>)}
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
  classes: PropTypes.object.isRequired,
  title: PropTypes.string.isRequired,
  titleize: PropTypes.bool,
  capitalize: PropTypes.bool,
  subtitle: PropTypes.string,
  icon: PropTypes.node,
};

PageTitle.defaultProps = {
  titleize: false,
  capitalize: false,
  subtitle: '',
  icon: null,
};

module.exports = withRouter(withStyles(style)(PageTitle));
