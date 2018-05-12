import React from 'react';
import { render } from 'react-dom';
import { BrowserRouter, Route } from 'react-router-dom';
import { createStore, applyMiddleware } from 'redux';
import thunk from 'redux-thunk';
import { Provider } from 'react-redux';
import { SpeciesShow } from './components/Species';
import { WordsShow } from './components/Words';
import { NamersShow } from './components/Namer';
import { CulturesShow } from './components/Culture';
import { TownForm, AreaShow } from './components/Area';
import { App, Home } from './components/App';
import rootReducer from './reducers';
import './main.scss';

const store = createStore(rootReducer, applyMiddleware(thunk));
render(
  (
    <Provider store={store}>
      <BrowserRouter>
        <App>
          <Route path="/home" component={Home} />
          <Route path="/" component={Home} />
          <Route path="/namers/:name" component={NamersShow} />
          <Route path="/words/:name" component={WordsShow} />
          <Route path="/species/:name" component={SpeciesShow} />
          <Route path="/cultures/:name" component={CulturesShow} />
          <Route path="/towns" component={TownForm} />
          <Route path="/towns/:id" component={AreaShow} />
        </App>
      </BrowserRouter>
    </Provider>
  ),
  document.getElementById('root'),
);
