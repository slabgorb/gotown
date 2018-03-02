import speciesApi from '../api/speciesApi';

export const loadSpeciesSuccess = (species) => ({ type:LOAD_SPECIES_SUCCESS, species })

export const loadSpecies = () => dispatch => speciesApi.loadSpecies()
  .then(species => dispatch(loadSpeciesSuccess(species)));

