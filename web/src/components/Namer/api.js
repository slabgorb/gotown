import axios from 'axios';

class namerApi {
  static getAll() {
    return axios.get('/api/namers').then(response => response.data);
  }
  static get(name) {
    return axios.get(`/api/namers/${name}`).then(response => response.data);
  }
}

export default namerApi;