<template>
  <section class="container">
    <h2>領収書印刷</h2>
    <ul>
      <li>
        <nuxt-link :to="{ name: 'receipts-id', params: { id: '1' } }">
          aaa
        </nuxt-link>
      </li>
      <li>user</li>
      <li>user</li>
    </ul>
  </section>
</template>

<script>
import Vue from "vue";

export default Vue.extend({
  data() {
    return {
      articles: [],
    };
  },
  mounted() {
    this.$firestore
      .collection('articles')
      .get()
      .then(articles => articles.forEach(article => {
        console.log(article)
      }))
  }
})
</script>
