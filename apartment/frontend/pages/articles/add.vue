<template>
  <Container>
    <Content>
      <Breadcrumb>
        <li><nuxt-link to="/articles">物件の管理</nuxt-link></li>
        <li>物件を追加する</li>
      </Breadcrumb>
    </Content>
    <Content>
      <Subheader title="物件を追加する">
        <button class="btn btn__primary">保存</button>
      </Subheader>
    </Content>
    <Content>
      <ArticleEditor v-model="article" :editable="true" />
    </Content>
  </Container>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import { Action } from "vuex-class";
import ArticleEditor from '~/components/ArticleEditor.vue';
import { ArticleDetail } from "~/declare/article";

@Component({
  data: () => ({
    article: {
      id: 1,
      name: "新しいマンション",
      rooms:[
        {
          id: 0,
          name: "101号室",
          rent: 100000,
          index: 0,
        }
      ],
    },
  }),
  components: {
    ArticleEditor,
  }
})
export default class extends Vue {
  @Action("setTitle", { namespace: "app" })
  setTitle;

  article: ArticleDetail | null = null;

  created() {
    this.setTitle("物件を追加する");
  }
}
</script>

