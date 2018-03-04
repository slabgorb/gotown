import * as types from '../types';

const initialState = { isFetching: false };

function appReducer(state = initialState, action) {
  switch (action.type) {
    case types.SELECT_SPECIES:
      return Object.assign({}, state, { isFetching: true });
    case types.RECEIVE_SPECIES:
      return Object.assign({}, state, {
        species: action.data,
      });
    default:
      return state;
  }
}

export default appReducer;
