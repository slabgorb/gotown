import React from 'react';
import Species from 'components/Species.jsx';
import Area from 'components/Area.jsx'
import Nav from 'components/Nav.jsx';
import { Route } from 'react-router-dom'
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';


class App extends React.Component {
  render() {
    return(
      <MuiThemeProvider>
        <div>

          <Nav/>
          <Route path="/species" component={Species}/>
          <Route path="/area" component={Area}/>
        </div>
      </MuiThemeProvider>
    )
  }
}

module.exports = App
