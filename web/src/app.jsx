import React from 'react';
import { render } from 'react-dom';
import { BrowserRouter } from 'react-router-dom';
import { createStore, applyMiddleware } from 'redux';
import thunk from 'redux-thunk';
import { Provider } from 'react-redux';
import SpeciesIndex from './components/Species/Index';
import App from './components/App';
import rootReducer from './reducers';
import routes from './routes';

const store = createStore(rootReducer, applyMiddleware(thunk));
render(
  (
    <Provider store={store}>
      <BrowserRouter routes={routes}>
        <App>
          <SpeciesIndex />
        </App>
      </BrowserRouter>
    </Provider>
  ),
  document.getElementById('root'),
);
