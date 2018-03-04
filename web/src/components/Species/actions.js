import axios from 'axios';
import * as types from '../../types';

export const setSpecies = name => ({ type: types.SELECT_SPECIES, name });
export const requestSpecies = name => ({ type: types.REQUEST_SPECIES, name });
export const receiveSpecies = (name, data) => ({ type: types.RECEIVE_SPECIES, name, data });

export const fetchSpecies = name => (dispatch) => {
  dispatch(requestSpecies(name));
  return axios.get(`data/${name}.json`)
    .then(res => dispatch(receiveSpecies(name, res.data)));
};
