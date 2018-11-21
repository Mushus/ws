<template>
  <Container>
    <Content>
      <Breadcrumb>
        <li>入居者の管理</li>
      </Breadcrumb>
    </Content>
    <Content>
      <Subheader title="入居者の管理" />
    </Content>
    <Content v-for="article in articles" :key="article.id">
      <h3 class="article-name">{{article.name}}</h3>
      <ul class="list">
        <li v-for="room in article.rooms" :key="room.id">
          <nuxt-link :to="`/tenants/roomstatus/${room.id}`">
            {{ room.name }}
            <span v-if="room.tenants.length > 0" class="tenant-name">({{ room.tenants[0].name }})</span>
            <span v-else class="badge-tag badge-tag__red">空き</span>
          </nuxt-link>
        </li>
      </ul>
    </Content>
  </Container>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator"
import { Action, Getter } from "vuex-class"
import { SearchTenantArticle } from '~/declare/article.ts';
@Component({
  data: () => ({
    articles: [],
  })
})
export default class extends Vue {
  @Action('setTitle', { namespace: 'app' }) setTitle
  articles: SearchTenantArticle[] = [];

  created() {
    this.setTitle('入居者の管理');
    this.updateArticles();
  }

  async updateArticles() {
    const data = await this.$axios.$get(`./api/articles/all.json`);
    this.articles = data;
  }
}
</script>

<style lang="scss" scoped>
.article-name {
  font-size: 1.7rem;
}
.list{
  border-top: 3px solid #eee;
  border-bottom: 2px solid #eee;
  li {
    border-top: 2px solid #eee;
    a {
      display: block;
      padding: 10px 20px;
      &:hover {
        background-color: #eee;
      }
    }
  }
}
</style>
