<template>
  <section>
    <b-breadcrumb>
      <b-breadcrumb-item :active="true">入居状況の確認</b-breadcrumb-item>
      <b-breadcrumb-item :active="true">建物を選択する</b-breadcrumb-item>
    </b-breadcrumb>
    <b-navbar>
      <h2>入居状況の確認</h2>
    </b-navbar>
    <b-container>
      <p>建物を選択してください。</p>
    </b-container>
    <b-list-group>
      <b-list-group-item
        v-for="article in articles"
        :key="article.id"
        :to="{ name:'occupants-articleId', params: { articleId: article.id } }"
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
      articlesDoc.forEach(article =>
        articles.push(normalizeArticle(article.id, article.data())));
    } catch (e) {
      console.log(e);
      return error({ statusCode: 500, message: 'データ取得失敗' });
    }

    return {
      articles,
    };
  }
})
</script>
