import axios from 'axios';

export const user_url = 'http://192.168.43.141:4000/api/user'
export default axios.create({
  baseURL: user_url
});
