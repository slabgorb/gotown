import { MuiThemeProvider, createMuiTheme, withStyles } from 'material-ui/styles';
import Reboot from 'material-ui/Reboot';
import React from 'react';
import PropTypes from 'prop-types';
import AppBar from 'material-ui/AppBar';
import Toolbar from 'material-ui/Toolbar';
import Button from 'material-ui/Button';
import Typography from 'material-ui/Typography';
import axios from 'axios';
import SimpleBottomNavigation from '../Nav/SimpleBottomNavigation';

const ralewayFF = ({ fontFamily: 'Raleway' });
const montserratFF = ({ fontFamily: 'Montserrat' });

const theme = createMuiTheme({
  palette: {
    primary: {
      light: '#757ce8',
      main: '#3f50b5',
      dark: '#002884',
      contrastText: '#fff',
    },
    secondary: {
      light: '#ff7961',
      main: '#f44336',
      dark: '#ba000d',
      contrastText: '#000',
    },


  },
  typography: {
    fontFamily: 'Raleway',
    body1: ralewayFF,
    body2: ralewayFF,
    display1: ralewayFF,
    display2: ralewayFF,
    display3: ralewayFF,
    display4: ralewayFF,
    caption: ralewayFF,
    headline: montserratFF,
    title: montserratFF,
  },
});

const seed = () => axios.put('/api/seed');

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
                <Button variant="raised" onClick={() => seed()}>Seed</Button>
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
