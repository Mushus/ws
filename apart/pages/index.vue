<template>
  <div>
    <h2>Hello world!</h2>
    <ul>
      <li v-for="article in articles" :key="article.id">{{ article.name }}</li>
    </ul>
  </div>
</template>

<script lang="ts">
import {
  Component,
  Vue
} from "nuxt-property-decorator"
import { State } from "vuex-class"
import Card from "~/components/Card.vue"

@Component({
  components: {
    Card
  }
})
export default class extends Vue {
  @State articles
}
</script>
<style scoped>
</style>
