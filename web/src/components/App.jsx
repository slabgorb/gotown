import { MuiThemeProvider, createMuiTheme, withStyles } from 'material-ui/styles';
import Reboot from 'material-ui/Reboot';
import React from 'react';
import PropTypes from 'prop-types';
import AppBar from 'material-ui/AppBar';
import Toolbar from 'material-ui/Toolbar';
import Typography from 'material-ui/Typography';
import SimpleBottomNavigation from './Nav/SimpleBottomNavigation';


const theme = createMuiTheme({
  palette: {
    primary: {
      main: '#37474f',
      light: '#62727b',
      dark: '#102027',
      contrastText: '#fff',
    },
    secondary: {
      main: '#afbdc4',
      light: '#e1eff7',
      dark: '#808d94',
      contrastText: '#000',
    },

  },
});

const App = ({ children, classes }) =>
  (
    <div>
      <Reboot>
        <MuiThemeProvider theme={theme}>
          <div>
            <AppBar position="static">
              <Toolbar>
                <Typography variant="title" color="inherit" className={classes.flex}>
                  Gotown
                </Typography>
              </Toolbar>
            </AppBar>
            {children}
            <SimpleBottomNavigation />
          </div>
        </MuiThemeProvider>
      </Reboot>
    </div>
  );

App.propTypes = {
  children: PropTypes.node.isRequired,
  classes: PropTypes.object.isRequired,
};

export default withStyles(theme)(App);
