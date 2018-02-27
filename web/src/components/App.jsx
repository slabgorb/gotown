import { MuiThemeProvider, createMuiTheme } from 'material-ui/styles';
import BottomNavigation, { BottomNavigationAction } from 'material-ui/BottomNavigation';
import Reboot from 'material-ui/Reboot';
import React from 'react';
import PropTypes from 'prop-types';
import { Fingerprint } from 'material-ui-icons/Fingerprint';
import { Face } from 'material-ui-icons/Face';
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

const App = ({ children }) =>
  (
    <div>
      <Reboot>
        <MuiThemeProvider theme={theme}>
          <div>
            {children}
            <SimpleBottomNavigation />
          </div>
        </MuiThemeProvider>
      </Reboot>
    </div>
  );

App.propTypes = {
  children: PropTypes.object.isRequired,
};

export default App;