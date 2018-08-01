import CssBaseline from '@material-ui/core/CssBaseline';
import React from 'react';
import { render } from 'react-dom';
import { Provider } from 'react-redux';
import { BrowserRouter, Route } from 'react-router-dom';
import { applyMiddleware, createStore } from 'redux';
import thunk from 'redux-thunk';
import { App, Home } from './components/App';
import { AreaShow, TownForm } from './components/Area';
import { BeingShow } from './components/Being';
import { CulturesShow } from './components/Culture';
import { NamersShow } from './components/Namer';
import { SpeciesShow } from './components/Species';
import { WordsShow } from './components/Words';
import './main.scss';
import rootReducer from './reducers';

const store = createStore(rootReducer, applyMiddleware(thunk));
render(
  (
    <React.Fragment>
      <CssBaseline />
      <Provider store={store}>
        <BrowserRouter>
          <App>
            <Route exact path="/home" component={Home} />
            <Route exact path="/" component={Home} />
            <Route exact path="/namers/:id" component={NamersShow} />
            <Route exact path="/words/:id" component={WordsShow} />
            <Route exact path="/species/:id" component={SpeciesShow} />
            <Route exact path="/cultures/:id" component={CulturesShow} />
            <Route exact path="/create/town" component={TownForm} />
            <Route exact path="/towns/:id" component={AreaShow} />
            <Route exact path="/beings/:id" component={BeingShow} />
          </App>
        </BrowserRouter>
      </Provider>
    </React.Fragment>
  ),
  document.getElementById('root'),
);
