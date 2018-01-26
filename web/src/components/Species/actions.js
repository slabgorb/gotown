import axios from 'axios';
import * as SpeciesTypes from './types';

export const setSpecies = name => ({ type: SpeciesTypes.SELECT_SPECIES, name });
export const requestSpecies = name => ({ type: SpeciesTypes.REQUEST_SPECIES, name });
export const receiveSpecies = (name, data) => ({ type: SpeciesTypes.RECEIVE_SPECIES, name, data });

export const fetchSpecies = name => (dispatch) => {
  dispatch(requestSpecies(name));
  return axios.get(`data/${name}.json`)
    .then(res => dispatch(receiveSpecies(name, res.data)));
};
