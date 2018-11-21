<template>
  <input class="input" type="date" v-model="model" :readonly="!editable">
</template>

<script lang="ts">
import * as numeral from 'numeral';
import Vue from "vue";
import { Component, Prop, Watch, Emit } from "nuxt-property-decorator";
import { ArticleDetail, Room } from '~/declare/article';
import * as moment from 'moment'

@Component({
  data: () => ({
    model: null,
    draggingRoom: null,
  }),
  props: {
    value: String,
    editable: {
      default: false,
      type: Boolean
    }
  }
})
export default class DateInput extends Vue {
  model: string | null = null;
  value: any;

  created () {
    this.onValueUpdate(this.value, null);
  }

  @Watch("value")
  onValueUpdate(val: any, oldVal: any) {
    if (val == "") {
      this.model = "";
    } else {
      this.model = moment(val).format("YYYY-MM-DD");
    }
  }

  @Watch("model", { deep: true })
  onModelUpdate(val: string, oldVal: any) {
    if (val == '') {
      this.$emit('input', "");
    } else {
      this.$emit('input', moment(val).format());
    }
  }
}
</script>

<style lang="scss" scoped>
</style>

