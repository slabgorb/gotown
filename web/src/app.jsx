import "main.scss"
import React from 'react';
import { render } from 'react-dom'
import { BrowserRouter } from 'react-router-dom'
import App from 'components/App.jsx'
import { dispatch, createStore } from 'redux';
import {Provider} from 'react-redux';
import { rootReducer } from './reducers.js'


let store = createStore(rootReducer)
render(
  (<Provider store={store}>
    <BrowserRouter>
      <App/>
    </BrowserRouter>
  </Provider>),
  document.getElementById('root')
)
