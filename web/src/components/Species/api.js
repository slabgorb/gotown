import axios from 'axios';

class speciesApi {
  static getAll() {
    return axios.get('/api/species').then(response => response.data).catch(error => error);
  }
  static get(id) {
    return axios.get(`/api/species/${id}`).then(response => response.data).catch(error => error);
  }
  static getExpression(id, genes) {
    return axios.get(`/api/species/${id}/expression`, { params: { chromosome: genes.join('|') } }).then(response => response.data).catch(error => error);
  }
}

export default speciesApi;
