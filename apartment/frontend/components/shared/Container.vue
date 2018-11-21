<template>
  <div class="container">
    <slot />
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import { Component } from "nuxt-property-decorator";

@Component({})
export default class Container extends Vue {
}
</script>

<style lang="scss" scoped>
.container {
  margin: 15px;
}
</style>
