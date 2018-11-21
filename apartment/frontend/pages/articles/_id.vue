<template>
  <Container>
    <Content>
      <Breadcrumb>
        <li><nuxt-link to="/articles">物件の管理</nuxt-link></li>
        <li>
          <template v-if="article">{{ article.name }}の詳細</template>
          <template v-else>物件の詳細</template>
        </li>
      </Breadcrumb>
    </Content>
    <Content v-if="article">
      <Subheader :title="`${article.name}の詳細`">
        <template v-if="editable && article">
          <button class="btn btn__primary" @click="editable = false">保存</button>
        </template>
        <template v-else>
          <button class="btn btn__primary" @click="editable = true">編集</button>
        </template>
      </Subheader>
    </Content>
    <Content>
      <ArticleEditor v-model="article" :editable="editable" />
    </Content>
  </Container>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator"
import { Action, Getter } from "vuex-class"
import ArticleEditor from '~/components/ArticleEditor.vue';
import { ArticleDetail } from "~/declare/article";

@Component({
  data: () => ({
    editable: false,
    article: null,
  }),
  components: {
    ArticleEditor,
  }
})
export default class extends Vue {
  article: ArticleDetail | null = null;
  @Action('setTitle', { namespace: 'app' }) setTitle

  created() {
    this.setTitle('物件の詳細')
    this.updateArticle(+this.$route.params.id);
  }

  async updateArticle(id: number) {
    const data = await this.$axios.$get(`./api/articles/${id}.json`);
    this.article = data;
  }
}
</script>

<style scoped>
</style>
