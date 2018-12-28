<template>
  <nav>
    <ul>
      <li><nuxt-link to="/articles">物件管理</nuxt-link></li>
    </ul>
  </nav>
</template>
<script lang="ts">
import {
  Component,
  Vue
} from "nuxt-property-decorator"

@Component({
  components: {}
})
export default class extends Vue {}
</script>
