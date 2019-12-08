import axios from 'axios';

export default axios.create({
  baseURL: 'http://192.168.43.15:3000/api'
});
