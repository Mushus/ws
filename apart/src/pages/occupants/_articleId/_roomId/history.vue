<template>
  <section>
    <b-breadcrumb>
      <b-breadcrumb-item :active="true">入居状況の確認</b-breadcrumb-item>
      <b-breadcrumb-item :to="{ name: 'occupants' }">建物を選択する</b-breadcrumb-item>
      <b-breadcrumb-item :to="{ name: 'occupants-articleId', params: { articleId } }">部屋を選択する</b-breadcrumb-item>
      <b-breadcrumb-item :active="true">入居する</b-breadcrumb-item>
    </b-breadcrumb>
    <b-navbar class="pt-3">
      <h2>入退去を行う</h2>
    </b-navbar>
  </section>
</template>

<script>
import Vue from 'vue';

export default Vue.extend({
  async asyncData({ params: { articleId, roomId } }) {
    return {
      articleId,
      roomId,
    };
  }
});
</script>
