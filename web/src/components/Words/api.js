import axios from 'axios';

class wordsApi {
  static getAll() {
    return axios.get('/api/words').then(response => response.data);
  }
  static get(name) {
    return axios.get(`/api/words/${name}`).then(response => response.data);
  }
}

export default wordsApi;