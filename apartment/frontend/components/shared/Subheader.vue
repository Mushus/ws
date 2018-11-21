<template>
  <header class="subheader">
    <h2>{{title}}</h2>
    <div class="control-group">
      <slot />
    </div>
  </header>
</template>

<script lang="ts">
import Vue from "vue";
import { Component } from "nuxt-property-decorator";

@Component({
  props: {
    title: {
      type: String,
      default: "",
    }
  }
})
export default class Subheader extends Vue {
}
</script>

<style lang="scss" scoped>
.subheader {
  display: flex;
  h2 {
    font-size: 20px;
  }
  .control-group {
    flex-grow: 1;
    display: flex;

    & > * {
      margin-left: 10px;
    }
    :first-child {
      margin-left: auto;
    }
  }
}
</style>
