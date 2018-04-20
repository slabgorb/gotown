import axios from 'axios';

class chromosomeApi {
  static random(count = 16) {
    return axios.get('/api/random/chromosome', { params: { count } }).then(resp => resp.data);
  }
}
export default chromosomeApi;
