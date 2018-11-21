<template>
  <div class="input-item">
    <slot />
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import { Component } from "nuxt-property-decorator";

@Component({})
export default class InputItem extends Vue {
}
</script>

<style lang="scss" scoped>
.input-item:not(:last-child) {
  margin-bottom: 10px;
}
</style>
