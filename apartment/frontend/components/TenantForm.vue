<template>
  <div>
    <InputItem>
      <label :for="`tenant${index}-name`">名前</label>
      <input class="input" :id="`tenant${index}-name`" type="text" v-model="model.name" :readonly="!editable">
    </InputItem>
    <InputItem>
      <label :for="`tenant${index}-rent`">家賃</label>
      <input class="input" :id="`tenant${index}-rent`" type="number" min="0" step="1000" v-model="model.rent" :readonly="!editable">
    </InputItem>
    <InputItem class="input-item__flex">
      <div class="input-item--group">
        <label :for="`tenant${index}-moveInAt`">入居日</label>
        <div class="form-group">
          <DateInput :id="`tenant${index}-moveInAt`" v-model="model.moveInAt" :editable="editable" />
          <button v-if="editable" class="btn btn__primary" @click="model.moveInAt = today">今日</button>
        </div>
      </div>
      <div class="input-item--group">
        <label :for="`tenant${index}-moveInAt`">退去日</label>
        <div class="form-group">
          <DateInput :id="`tenant${index}-moveOutAt`" v-model="model.moveOutAt" :editable="editable" />
          <button v-if="editable" class="btn btn__primary" @click="model.moveOutAt = today">今日</button>
        </div>
      </div>
    </InputItem>
  </div>
</template>

<script lang="ts">
import * as numeral from 'numeral';
import * as moment from 'moment';
import Vue from "vue";
import { Component, Prop, Watch, Emit } from "nuxt-property-decorator";
import { Tenant } from '~/declare/article';

@Component({
  data: () => ({
    index: Math.random().toString(36).slice(-8),
    model: null,
  }),
  props: {
    value: Object,
    editable: {
      default: false,
      type: Boolean
    }
  }
})
export default class ArticleEditor extends Vue {
  model: Tenant | null = null;
  value: any;

  created () {
    this.model = this.value;
  }

  @Watch("value")
  onValueUpdate(val: any, oldVal: any) {
    this.model = val;
  }

  @Watch("model", { deep: true })
  onModelUpdate(val: any, oldVal: any) {
    this.$emit('input', val)
  }


  get today() {
    const now = new Date();
    return moment({
      y: now.getFullYear(),
      M: now.getMonth(),
      d: now.getDate(),
      h: 0,
      m: 0,
      s: 0,
      ms: 0,
    }).utcOffset(now.getTimezoneOffset()).toISOString();
  }

  get forever() {
    const now = new Date();
    return moment({
      y: 2999,
      M: 1,
      d: 1,
      h: 0,
      m: 0,
      s: 0,
      ms: 0,
    }).utcOffset(now.getTimezoneOffset()).toISOString();
  }
}
</script>

<style lang="scss" scoped>
</style>

