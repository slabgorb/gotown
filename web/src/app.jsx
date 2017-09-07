import "main.scss"
import 'typeface-roboto'
import React from 'react';
import ReactDOM from 'react-dom'
//import { MuiThemeProvider, createMuiTheme } from 'material-ui/styles/MuiThemeProvider';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider'
import Species from 'Species.jsx'

//var theme = function()  {
  //return createMuiTheme({
    // typography: {
    //   fontFamily: "'Playfair Display', serif",
    //   headline: {
    //     fontFamily: "'Josefin Slab', serif",
    //   }
    // }
  //});
//}

ReactDOM.render(
  <MuiThemeProvider>
      <Species name='viking'/>
  </MuiThemeProvider>,
  document.getElementById('root')
)
