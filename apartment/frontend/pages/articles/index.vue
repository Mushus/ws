<template>
  <Container>
    <Content>
      <Breadcrumb>
        <li>物件の管理</li>
      </Breadcrumb>
    </Content>
    <Content>
      <Subheader title="物件の管理">
        <button class="btn btn__primary" @click="$router.push('/articles/add')">追加</button>
      </Subheader>
    </Content>
    <Content>
      <ul class="list">
        <li v-for="article in articles" :key="article.id">
          <nuxt-link :to="`/articles/${article.id}`">{{ article.name }}</nuxt-link>
        </li>
      </ul>
    </Content>
  </Container>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator"
import { Action, Getter } from "vuex-class"
import { Article } from '~/declare/article.ts';
@Component({
  data: () => ({
    articles: [],
  })
})
export default class extends Vue {
  @Action('setTitle', { namespace: 'app' }) setTitle
  articles: Article[] = [];

  created() {
    this.setTitle('物件の管理');
    this.updateArticles();
  }

  async updateArticles() {
    const data = await this.$axios.$get(`./api/articles.json`);
    this.articles = data;
  }
}
</script>

<style lang="scss" scoped>
.list{
  border-bottom: 2px solid #eee;
  li {
    border-top: 2px solid #eee;
    a {
      display: block;
      padding: 10px;
      &:hover {
        background-color: #eee;
      }
    }
  }
}
</style>
