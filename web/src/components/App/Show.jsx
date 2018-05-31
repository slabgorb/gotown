import { MuiThemeProvider, createMuiTheme, withStyles } from '@material-ui/core/styles';
import PropTypes from 'prop-types';
import React from 'react';


const ralewayFF = ({ fontFamily: 'Raleway' });
const montserratFF = ({ fontFamily: 'Montserrat' });

const styles = () => ({
  flex: { flex: 1 },
  main: {  },
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
