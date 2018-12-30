<template>
  <section>
    <h2>物件一覧</h2>
    <section>
      <nuxt-link :to="{ name: 'articles-create' }">物件を作成する</nuxt-link>
    </section>
    <ul>
      <li v-for="article in articles" :key="article.id">
        <nuxt-link :to="{ name:'articles-id', params: { id: article.id } }">
          {{ article.data.name }}
        </nuxt-link>
      </li>
    </ul>
  </section>
</template>

<script>
import Vue from 'vue';

export default Vue.extend({
  async asyncData({ error, $firestore }) {
    const articlesRef = $firestore.collection('articles');

    const articles = [];
    try {
      const articlesDoc = await articlesRef.get();
      articlesDoc.forEach(article => articles.push({
        id: article.id,
        data: article.data()
      }));
    } catch (e) {
      console.log(e);
      return error({ statusCode: 500, message: 'データ取得失敗' });
    }

    return {
      articles: articles
    };
  },
  methods: {
  }
})
</script>
