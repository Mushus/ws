<template>
  <nav class="breadcrumb">
    <ul>
      <li><nuxt-link to="/"><i class="material-icons md-18">home</i></nuxt-link></li>
      <slot />
    </ul>
  </nav>
</template>

<script lang="ts">
import Vue from "vue";
import { Component } from "nuxt-property-decorator";

@Component({})
export default class ArticleEditor extends Vue {
}
</script>

<style lang="scss" scoped>
.breadcrumb {
  li {
    display: inline;
    color: #888;

    &:not(:last-child)::after {
      margin: 10px;
      content: ">";
      color: #666;
    }
  }
}
</style>
