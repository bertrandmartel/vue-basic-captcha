import login from '../../api/login'

const state = {
  showPassword: false,
  formValid: false,
  password: '',
  username: '',
  usernameRules: [
    v => !!v || 'Username is required',
  ],
  passwordRules: [
    v => !!v || 'Password is required',
  ],
}

const mutations = {
  updateShowPassword(state, value){
    state.showPassword = value
  },
  updatePassword(state, value){
    state.password = value
  },
  updateUsername(state, value){
    state.username = value
  },
  updateFormValid(state, value){
    state.formValid = value
  },
}

const actions = {
  login({state, rootState, commit, dispatch}, noFirstSelect){
    var captchaResponse = document.getElementById("g-recaptcha-response").value;
    if (!captchaResponse){
      return dispatch('core/dispatchMessage', {color: 'error', message: 'captcha is required'}, {root:true})
    }
    const formData = new FormData();
    formData.append('login', state.username);
    formData.append('password', state.password);
    formData.append('g-recaptcha-response', captchaResponse);
    login.login({
      serverPath: rootState.core.serverPath,
      formData: formData
    },
    (response) => {
      document.location.href = "/"
    },
    (err) => {
      console.log(err);
      if (err.response) {
        /*eslint-disable */
        if (grecaptcha){
          grecaptcha.reset();
        }
        /* eslint-enable */
        return dispatch('core/dispatchMessage', {color: 'error', message: "authentication failed"}, {root:true})
      }
    })
  }
}

export default {
    namespaced: true,
    state,
    mutations,
    actions
}