import axios from 'axios';

class cultureApi {
  static getAll() {
    return axios.get('/api/cultures').then(response => response.data);
  }
  static get(id) {
    return axios.get(`/api/cultures/${id}`).then(response => response.data);
  }
}

export default cultureApi;
