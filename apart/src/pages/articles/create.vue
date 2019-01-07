<template>
  <section>
    <b-breadcrumb>
      <b-breadcrumb-item :to="{ name: 'articles' }">物件一覧</b-breadcrumb-item>
      <b-breadcrumb-item :active="true">物件作成</b-breadcrumb-item>
    </b-breadcrumb>
    <b-navbar>
      <h2>物件作成</h2>
      <b-button-toolbar>
        <b-button-group>
          <b-button :to="{ name: 'articles' }">一覧へ</b-button>
        </b-button-group>
      </b-button-toolbar>
    </b-navbar>
    <b-form @submit.prevent="submit()">
      <b-form-group label="建物名">
        <b-form-input type="text" v-model="article.data.name" />
      </b-form-group>
      <b-button-group>
        <b-button variant="primary" type="submit">作成する</b-button>
      </b-button-group>
    </b-form>
  </section>
</template>

<script>
import Vue from 'vue';
import { normalizeArticle } from '@/util/normalize';

export default Vue.extend({
  data() {
    return {
      article: normalizeArticle(),
    };
  },
  methods: {
    async submit() {
      const articlesRef = this.$firestore.collection('articles');
      try {
        const article = await articlesRef.add(this.article.data);
        this.$router.push({
          name: 'articles-id',
          params: {
            id: article.id,
          },
        });
      } catch(e) {
        console.log(e);
      }
    }
  }
})
</script>

