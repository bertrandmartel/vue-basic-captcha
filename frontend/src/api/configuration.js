import axios from 'axios'

export default {
  getConfiguration(data, cb, errorCb) {
    axios.get(`${data.serverPath}/configuration`)
      .then(function(response){
        cb(response)
      })
      .catch(function(err){
        errorCb(err)
      });
  },
}