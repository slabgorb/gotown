export const SELECT_SPECIES = 'SELECT_SPECIES'
export const RECEIVE_SPECIES = 'RECEIVE_SPECIES'
export const REQUEST_SPECIES = 'REQUEST_SPECIES'

export function setSpecies(name) {
  return { type: SELECT_SPECIES, name }
}

export function requestSpecies(name) {
  return { type: REQUEST_SPECIES, name }
}

export function receiveSpecies(name, json) {
  return {
    type: RECEIVE_SPECIES,
    name,
    data: json.data
  }
}
