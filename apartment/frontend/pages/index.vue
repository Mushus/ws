<template>
  <Container>
    <h1 class="header">賃貸管理システム</h1>
    <div class="cards">
      開発中…
    </div>
  </Container>
</template>

<script lang="ts">
import {
  Component,
  Vue
} from "nuxt-property-decorator"
import { Action } from "vuex-class"

@Component({})
export default class extends Vue {
  @Action('setTitle', { namespace: 'app' }) setTitle

  created() {
    this.setTitle('賃貸管理システム')
  }
}
</script>

