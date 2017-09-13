import "main.scss"
import React from 'react';
import { render } from 'react-dom'
import { BrowserRouter } from 'react-router-dom'
//import { MuiThemeProvider, createMuiTheme } from 'material-ui/styles/MuiThemeProvider';
import App from 'components/App.jsx'
import { dispatch } from 'redux';

const defaultSpecies = {
  name: "",
  genetics: {},
  genderNames: {},
}

let store = createStore()

render(
  <Provider store={store}>
    <BrowserRouter>
      <App/>
    </BrowserRouter>
  </Provider>,
  document.getElementById('root')
)
