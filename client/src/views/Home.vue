<template>
  <v-container>
    <v-row>
      <img alt="Vue logo" src="../assets/logo.png" />
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

@Component({
  name: 'Home',
})
export default class Home extends Vue {}
</script>
