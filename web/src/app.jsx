import React from 'react';
import { render } from 'react-dom';
import { BrowserRouter, Route } from 'react-router-dom';
import { createStore, applyMiddleware } from 'redux';
import thunk from 'redux-thunk';
import { Provider } from 'react-redux';
import { SpeciesShow, SpeciesList } from './components/Species';
import { WordsShow, WordsList } from './components/Words';
import { CulturesShow, CulturesList } from './components/Culture';
import { TownForm, AreaList, AreaShow } from './components/Area';
import { App } from './components/App';
import rootReducer from './reducers';
import './main.scss';

const store = createStore(rootReducer, applyMiddleware(thunk));
render(
  (
    <Provider store={store}>
      <BrowserRouter>
        <App>
          <Route path="/words" component={WordsList} />
          <Route path="/words/:name" component={WordsShow} />
          <Route path="/species" component={SpeciesList} />
          <Route path="/species/:name" component={SpeciesShow} />
          <Route path="/cultures" component={CulturesList} />
          <Route path="/cultures/:name" component={CulturesShow} />
          <Route path="/towns" component={AreaList} />
          <Route path="/towns" component={TownForm} />
          <Route path="/towns/:name" component={AreaShow} />
        </App>
      </BrowserRouter>
    </Provider>
  ),
  document.getElementById('root'),
);
