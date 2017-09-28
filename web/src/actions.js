import axios from 'axios';

export const SELECT_SPECIES = 'SELECT_SPECIES'
export const RECEIVE_SPECIES = 'RECEIVE_SPECIES'
export const REQUEST_SPECIES = 'REQUEST_SPECIES'

export const setSpecies = (name) => ({ type: SELECT_SPECIES, name })
export const requestSpecies = (name) => ({ type: REQUEST_SPECIES, name })
export const receiveSpecies = (name, json) => ({ type: RECEIVE_SPECIES, name, data: json.data })

export const fetchSpecies = (name) => dispatch => {
  dispatch(requestSpecies(name))
  return axios.get(`data/${name}.json`)
    .then(res => {dispatch(receiveSpecies(name, res))})
}
