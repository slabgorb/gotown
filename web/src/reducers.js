import {
  SELECT_SPECIES,
  RECEIVE_SPECIES,
  setSpecies
} from './actions'
import { combineReducers } from 'redux'
const defaultSpecies = {
  name: "",
  genetics: {},
  genderNames: {},
}

const defaultArea = {
  name: "",
  residents: [],
}
function speciesReducer(state = {isFetching: false, species: defaultSpecies, area: defaultArea}, action) {
  switch (action.type) {
    case SELECT_SPECIES:
      return Object.assign({}, state, {isFetching: true})
    case RECEIVE_SPECIES:
      return Object.assign({}, state, {
        species: action.data
      })

    default:
      return state
  }
}


export const rootReducer = combineReducers( {
  speciesReducer
})
