import axios from 'axios';

export const SELECT_SPECIES = 'SELECT_SPECIES'
export const RECEIVE_SPECIES = 'RECEIVE_SPECIES'
export const REQUEST_SPECIES = 'REQUEST_SPECIES'

export function setSpecies(name) {
  return { type: SELECT_SPECIES, name }
}

export function requestSpecies(name) {
  return { type: REQUEST_SPECIES, name }
}

function receiveSpecies(name, json) {
  return {
    type: RECEIVE_SPECIES,
    name,
    data: json.data
  }
}

function fetchSpecies(name) {
  return dispatch => {
    dispatch(requestSpecies(name))
    return axios.get(`data/${this.props.name}.json`)
      .then(res => {dispatch(receiveSpecies(name, res))})
  }
}
