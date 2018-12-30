<template>
  <section>
    <h2>物件作成</h2>
    <form @submit="e => (submit(), false)">
      <section>
        <dl>
          <dt><label for="article_name">物件名</label></dt>
          <dd><input id="article_name" type="text" v-model="article.data.name" /></dd>
        </dl>
        <button type="submit">作成する</button>
      </section>
    </form>
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
        return error({ statusCode: 500, message: 'データ取得失敗' });
      }
    }
  }
})
</script>

