import axios from 'axios';

class heraldryApi {
  static random() {
    return axios.get('/api/random/heraldry').then(resp => resp.data);
  }
}

export default heraldryApi;
