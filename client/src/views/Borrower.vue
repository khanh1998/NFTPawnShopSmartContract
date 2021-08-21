<template>
  <v-container>
    <v-row>
      <v-col>
        <p>Hello borrower!!!</p>
      </v-col>
    </v-row>
    <v-row>
      <pawn-creator/>
    </v-row>
  </v-container>
</template>
<script lang="ts">
import { Vue, Component } from 'vue-property-decorator';
import PawnCreator from '@/components/PawnCreator.vue';

@Component({
  components: { PawnCreator },
  name: 'Borrower',
})
export default class extends Vue {
}
</script>
