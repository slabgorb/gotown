import "main.scss"
import React from 'react';
import ReactDOM from 'react-dom'
import { BrowserRouter } from 'react-router-dom'
//import { MuiThemeProvider, createMuiTheme } from 'material-ui/styles/MuiThemeProvider';
import App from 'components/App.jsx'

ReactDOM.render(
  (<BrowserRouter>
    <App/>
  </BrowserRouter>),
  document.getElementById('root')
)
