import axios from 'axios';

class wordsApi {
  static getAll() {
    return axios.get('/api/words').then(response => response.data);
  }
  static get(name, id) {
    return axios.get(`/api/words/${name || id}`).then(response => response.data);
  }
}

export default wordsApi;