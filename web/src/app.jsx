import "main.scss"
import 'typeface-roboto'
import React from 'react';
import ReactDOM from 'react-dom'
var Area = require('area.jsx').Area
var AreaForm = require('area.jsx').AreaForm
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
ReactDOM.render(
  <MuiThemeProvider>
    <div>

      <AreaForm/>
      <Area/>
    </div>
  </MuiThemeProvider>,
  document.getElementById('root')
)
