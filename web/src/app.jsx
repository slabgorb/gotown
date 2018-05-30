import React from 'react';
import { render } from 'react-dom';
import { Provider } from 'react-redux';
import { BrowserRouter, Route } from 'react-router-dom';
import { applyMiddleware, createStore } from 'redux';
import thunk from 'redux-thunk';
import { App, Home } from './components/App';
import { AreaShow, TownForm } from './components/Area';
import { CulturesShow } from './components/Culture';
import { NamersShow } from './components/Namer';
import { SpeciesShow } from './components/Species';
import { WordsShow } from './components/Words';
import './main.scss';
import rootReducer from './reducers';

const store = createStore(rootReducer, applyMiddleware(thunk));
render(
  (
    <Provider store={store}>
      <BrowserRouter>
        <App>
          <Route exact path="/home" component={Home} />
          <Route exact path="/" component={Home} />
          <Route path="/namers/:name" component={NamersShow} />
          <Route path="/words/:name" component={WordsShow} />
          <Route path="/species/:name" component={SpeciesShow} />
          <Route path="/cultures/:name" component={CulturesShow} />
          <Route exact path="/towns/create" component={TownForm} />
          <Route path="/towns/:id" component={AreaShow} />
        </App>
      </BrowserRouter>
    </Provider>
  ),
  document.getElementById('root'),
);
