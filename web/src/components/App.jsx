import { MuiThemeProvider, createMuiTheme } from 'material-ui/styles';
import Reboot from 'material-ui/Reboot';
import React from 'react';
import Species from './Species/index';


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
      contrastText:'#000',
    }

  }
});

const App = () =>
  (
    <div>
      <Reboot>
        <MuiThemeProvider theme={theme}>
          <Species name="human" />
        </MuiThemeProvider>
      </Reboot>
    </div>
  );
module.exports = App;
