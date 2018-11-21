<template>
  <section class="content">
    <slot />
  </section>
</template>

<script lang="ts">
import Vue from "vue";
import { Component } from "nuxt-property-decorator";

@Component({})
export default class Content extends Vue {
}
</script>

<style lang="scss" scoped>
.content {
  margin-bottom: 20px;
}
</style>
