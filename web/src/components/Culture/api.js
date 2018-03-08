import axios from 'axios';

class cultureApi {
  static getAll() {
    return axios.get('/api/cultures').then(response => response.data);
  }
  static get(name) {
    return axios.get(`/api/cultures/${name}`).then(response => response.data);
  }
}

export default cultureApi;
