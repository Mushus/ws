<template>
  <section>
    <b-navbar class="pt-3">
      <h2>物件作成</h2>
      <b-button-toolbar>
        <b-button-group>
          <b-button :to="{ name: 'articles-create' }">一覧へ</b-button>
        </b-button-group>
      </b-button-toolbar>
    </b-navbar>
    <b-form @submit="e => (submit(), false)">
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
import Vue from "vue"

export default Vue.extend({
  data() {
    return {
      article: {
        id: null,
        data: {
          name: '',
        },
      }
    }
  },
  methods: {
    async submit() {
      const articlesRef = this.$firestore.collection('articles');
      try {
        const article = await articlesRef.add(this.article.data);
        console.log(article)
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

