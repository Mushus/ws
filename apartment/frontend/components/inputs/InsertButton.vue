<template>
  <button class="button" @click="handleClick">
    <i class="material-icons">add</i>
  </button>
</template>

<script lang="ts">
import Vue from "vue";
import { Component, Prop, Watch, Emit } from "nuxt-property-decorator";

@Component({})
export default class InsertButton extends Vue {
  handleClick(e) {
    this.$emit('click', e);
  }
}
</script>

<style lang="scss" scoped>
.button {
  border: 2px dashed #2196F3;
  color: #2196F3;
  background-color: transparent;
  width: 100%;
  height: 40px;
  margin-bottom: 15px;
  &:hover {
    background-color: #E3F2FD;
  }
}
</style>

