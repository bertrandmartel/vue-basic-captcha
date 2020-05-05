<template>
  <v-content>
    <div>
        <pre>{{JSON.stringify(message, null, 2)}}</pre>
    </div>
  </v-content>
</template>

<script>
import { mapState } from 'vuex'

export default {
  name: 'Root',
  components: {
  },
  computed: {
    ...mapState({
      message: state => state.main.message,
    }),
  },
  created(){
    this.$store.dispatch('main/init');
  },
}
</script>