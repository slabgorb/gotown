import axios from 'axios';

class cultureApi {
  static getAll() {
    axios.get('/api/cultures').then(response => response.data).catch(error => error);
  }
  static get(name) {
    axios.get(`/api/cultures/${name}`).then(response => response.data).catch(error => error);
  }
}

export default cultureApi;
