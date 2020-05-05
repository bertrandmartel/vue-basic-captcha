import axios from 'axios'

export default {
  login(data, cb, errorCb) {
    axios.post(`${data.serverPath}/login`, data.formData)
      .then(function(response){
        cb(response)
      })
      .catch(function(err){
        errorCb(err)
      });
  },
  userinfo(data, cb, errorCb) {
    axios.get(`${data.serverPath}/userinfo`)
      .then(function(response){
        cb(response)
      })
      .catch(function(err){
        errorCb(err)
      });
  },
}