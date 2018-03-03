import React from 'react';
import { render } from 'react-dom';
import { BrowserRouter, Route } from 'react-router-dom';
import { createStore, applyMiddleware } from 'redux';
import thunk from 'redux-thunk';
import { Provider } from 'react-redux';
import { SpeciesShow, SpeciesList } from './components/Species';
import App from './components/App';
import rootReducer from './reducers';
import './main.scss';

const store = createStore(rootReducer, applyMiddleware(thunk));
render(
  (
    <Provider store={store}>
      <BrowserRouter>
        <App>
          <Route path="/species" component={SpeciesList} />
          <Route path="/species/:name" component={SpeciesShow} />
        </App>
      </BrowserRouter>
    </Provider>
  ),
  document.getElementById('root'),
);
