import axios from 'axios';

class speciesApi {
  static getAll() {
    return axios.get('/api/species').then(response => response.data).catch(error => error);
  }
  static get(name) {
    return axios.get(`/api/species/${name}`).then(response => response.data).catch(error => error);
  }
}

export default speciesApi;
