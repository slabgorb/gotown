import { createMuiTheme, MuiThemeProvider, withStyles } from '@material-ui/core/styles';
import PropTypes from 'prop-types';
import React from 'react';


const ralewayFF = (size, weight) => ({ fontFamily: 'Raleway', fontSize: size, fontWeight: weight });

const styles = () => ({
  flex: { flex: 1 },
  main: {},
  appBar: {},
  button: {
    color: 'white',
  },
});

const theme = createMuiTheme({
  palette: {
    primary: {
      light: '#757ce8',
      main: '#3f50b5',
      dark: '#002884',
      contrastText: '#fff',
    },
    secondary: {
      light: '#dd7961',
      main: '#454336',
      dark: '#34000d',
      contrastText: '#000',
    },


  },
  typography: {
    fontFamily: 'Raleway',
    body1: ralewayFF(16, 400),
    body2: ralewayFF(18, 600),
    display1: ralewayFF(32, 500),
    display2: ralewayFF(28, 500),
    display3: ralewayFF(24, 500),
    display4: ralewayFF(20, 500),
    caption: ralewayFF(14, 500),
    headline: ralewayFF(22, 600),
    title: ralewayFF(24, 500),
  },
});

const App = ({ children, classes }) => (
  <div>
    <MuiThemeProvider theme={theme}>
      <div className={classes.main}>
        {children}
      </div>
    </MuiThemeProvider>
  </div>
);

App.propTypes = {
  children: PropTypes.node.isRequired,
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(App);
