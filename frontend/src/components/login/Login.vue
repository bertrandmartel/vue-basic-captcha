<template>
  <v-container fluid>
    <v-card width="400" class="mx-auto mt-5 loginCard">
      <v-card-title class="pb-0" style='text-align: center; display: block;'>
        <h1>Login</h1>
      </v-card-title>
      <v-card-text>
        <v-form style='margin-top:20px;' v-model="valid">
          <v-text-field 
            required
            label="Username"
            autocorrect="off"
            autocapitalize="none"
            v-model="username"
            :rules="usernameRules"
            prepend-icon="mdi-account-circle"
          />
          <v-text-field 
            required
            :type="showPassword ? 'text' : 'password'" 
            label="Password"
            v-model="password"
            :rules="passwordRules"
            prepend-icon="mdi-lock"
            :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
            @click:append="showPassword = !showPassword"
          />
          <div class="captcha-container">
            <div class="g-recaptcha" data-sitekey="6LfPhvEUAAAAAGha5axkdtZCWnJu54XaBlgYSKeV"></div>
          </div>
        </v-form>
      </v-card-text>
      <v-divider></v-divider>
      <v-card-actions>
        <v-btn :disabled="!valid" color="primary" style='width:100%;' @click="login">Login</v-btn>
      </v-card-actions>
    </v-card>
  </v-container>
</template>

<script>
import { mapState } from 'vuex'

export default {
  name: 'Login',
  computed: {
    ...mapState({
      usernameRules: state => state.login.usernameRules,
      passwordRules: state => state.login.passwordRules,
    }),
    showPassword: {
      get() {
        return this.$store.state.login.showPassword
      },
      set(value) {
        this.$store.commit('login/updateShowPassword', value)
      }
    },
    password: {
      get() {
        return this.$store.state.login.password
      },
      set(value) {
        this.$store.commit('login/updatePassword', value)
      }
    },
    username: {
      get() {
        return this.$store.state.login.username
      },
      set(value) {
        this.$store.commit('login/updateUsername', value)
      }
    },
    valid: {
      get() {
        return this.$store.state.login.formValid
      },
      set(value) {
        this.$store.commit('login/updateFormValid', value)
      }
    },
  },
  methods: {
    login(){
      this.$store.dispatch('login/login')
    }
  },
}
</script>

<style>
@media screen and (min-height: 500px) {
  /*when not landscape center horizontally and vertically */
  .loginCard {
    position: absolute;
    left: 50%;
    top: 40%;
    -webkit-transform: translate(-50%, -50%);
    transform: translate(-50%, -50%);
  }
}
@media screen and (max-height: 499px) {
  /*when landscape add few space below */
  .loginCard {
    margin-bottom: 50px;
  }
}

@media screen and (max-width: 450px) {
  /*when mobile add margin left & right*/
  .loginCard {
    width: 350px !important;
  }
}

@media screen and (max-height: 500px) {
  /*when landscape hide the appbar*/
  .v-content {
    padding-top: 0 !important;
  }
}

.captcha-container {
    text-align: center;
    margin-top:10px;
}

.g-recaptcha {
    display: inline-block;
}
</style>

