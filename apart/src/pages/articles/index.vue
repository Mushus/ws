<template>
  <section>
    <b-breadcrumb>
      <b-breadcrumb-item :active="true">物件一覧</b-breadcrumb-item>
    </b-breadcrumb>
    <b-navbar>
      <h2>物件一覧</h2>
      <b-button-toolbar>
        <b-button-group>
          <b-button :to="{ name: 'articles-create' }" variant="primary">物件を作成する</b-button>
        </b-button-group>
      </b-button-toolbar>
    </b-navbar>
    <b-list-group class="pt-3">
      <b-list-group-item
        v-for="article in articles"
        :key="article.id"
        :to="{ name:'articles-id', params: { id: article.id } }"
        >
        {{ article.data.name }}
      </b-list-group-item>
    </b-list-group>
  </section>
</template>

<script>
import Vue from 'vue';
import { normalizeArticle } from '@/util/normalize';

export default Vue.extend({
  async asyncData({ error, $firestore }) {
    const articlesRef = $firestore.collection('articles');

    const articles = [];
    try {
      const articlesDoc = await articlesRef.get();
      articlesDoc.forEach(article => articles.push(normalizeArticle(article.id, article.data())));
    } catch (e) {
      console.log(e);
      return error({ statusCode: 500, message: 'データ取得失敗' });
    }

    return {
      articles: articles
    };
  },
})
</script>
