import axios from 'axios';
import cultureApi from '../Culture/api';
import speciesApi from '../Species/api';

class areaApi {
  static getSpecies() {
    return speciesApi.getAll();
  }
  static getCultures() {
    return cultureApi.getAll();
  }
  static get(name) {
    return axios.get(`/api/town/${name}`).then(resp => resp.data);
  }
}

export default areaApi;
