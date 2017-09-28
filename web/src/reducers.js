import {
  SELECT_SPECIES,
  RECEIVE_SPECIES,
  setSpecies
} from './actions'
import { combineReducers } from 'redux'
const defaultSpecies = {
  name: "",
  genetics: {},
}

const defaultArea = {
  name: "",
  residents: [],
}

const initialState = {isFetching: false, species: defaultSpecies, area: defaultArea}

export default function (state = initialState, action) {
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
