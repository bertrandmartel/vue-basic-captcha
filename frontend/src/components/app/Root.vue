<template>
  <div>
    <div class="loadProgress" v-if="!loaded">
      <v-progress-circular
        :size="70"
        indeterminate
        color="primary"
      ></v-progress-circular>
    </div>
    <div v-else class="root">
        <pre>{{JSON.stringify(message, null, 2)}}</pre>
    </div>
  </div>
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
      loaded: state => state.main.loaded,
    }),
  },
  created(){
    this.$store.dispatch('main/init');
  },
}
</script>

<style>
/*when not landscape center horizontally and vertically */
.loadProgress {
  position: absolute;
  left: 50%;
  top: 50%;
  -webkit-transform: translate(-50%, -50%);
  transform: translate(-50%, -50%);
}

@media screen and (max-height: 499px) {
  /*when landscape margin left */
  .root {
    margin-left: 60px;
  }
}

@media screen and (min-height: 500px) {
  .root {
    margin-top: 64px;
  }
}

.root {
  overflow-x: visible
}

</style>