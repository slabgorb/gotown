import axios from 'axios';

class beingApi {
  static get(id) {
    return axios.get(`/api/beings/${id}`).then(response => response.data);
  }
}

export default beingApi;
