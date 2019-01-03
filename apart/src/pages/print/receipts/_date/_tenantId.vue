<template>
  <section>
    <b-breadcrumb>
      <b-breadcrumb-item :to="{ name: 'print' }">書類を印刷</b-breadcrumb-item>
      <b-breadcrumb-item :active="true">領収書を印刷</b-breadcrumb-item>
    </b-breadcrumb>
    <b-navbar>
      <h2>領収書を印刷</h2>
      <b-button-toolbar>
        <b-button-group>
          <b-button @click="print()" variant="primary">印刷する</b-button>
        </b-button-group>
      </b-button-toolbar>
    </b-navbar>
    <iframe ref="printfield"></iframe>
  </section>
</template>

<script>
import Vue from 'vue';

export default Vue.extend({
  async asyncData({ error, params: { date } }) {
  },
  methods: {
    print() {
      this.$refs.printfield.contentWindow.print();
    }
  }
})
</script>
