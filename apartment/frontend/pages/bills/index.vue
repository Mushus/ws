<template>
  <Container>
    <Content>
      <Breadcrumb>
        <li>請求書の発行</li>
      </Breadcrumb>
    </Content>
    <Content>
      <Subheader title="請求書の発行" />
    </Content>
    <template v-if="result != null">
      <Content v-for="article in result.articles" :key="article.id">
        <table class="article-list">
          <caption>{{article.name}}</caption>
          <thead class="article-list--header">
            <tr>
              <th class="room-name-header">物件名</th>
              <th class="tenant-name-header">入居者名</th>
              <th class="issue-flag-header">発行済</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="room in article.rooms" :key="room.id" @click="$router.push(`/bills/tenants/${room.tenant.id}`)">
              <td>{{room.name}}</td>
              <td>{{room.tenant.name}}</td>
              <td class="check-column"><i class="material-icons md-18" v-if="room.tenant.issueFlag">check</i></td>
            </tr>
          </tbody>
        </table>
      </Content>
    </template>
  </Container>
</template>

<script lang="ts">
import {
  Component,
  Vue
} from "nuxt-property-decorator"
import { Action } from "vuex-class"
import { ListBill } from "~/declare/article";

@Component({
  data: () => ({
    result: null,
  })
})
export default class extends Vue {
  result: ListBill | null = null;
  @Action('setTitle', { namespace: 'app' }) setTitle

  created() {
    this.setTitle('請求書発行')
    this.updateArticles();
  }

  async updateArticles() {
    const data = await this.$axios.$get(`./api/bills/list.json`);
    this.result = data;
  }
}
</script>

<style lang="scss">
.article-list {
  width: 100%;
  box-sizing: border-box;
  table-layout: fixed;

  td,
  th {
    padding: 5px 10px;
  }
  caption {
    font-size: 1.7rem;
    text-align: left;
  }
  thead {
    border-bottom: 3px solid #ddd;
  }
  tbody {
    tr:nth-child(2n) {
      background-color: #f8f8f8;
    }
    tr:hover {
      background-color: #eee;
      cursor: pointer;
    }
  }
}
.tenant-name-header,
.room-name-header {
  width: 50;
}
.article-list--header {
  background-color: #eee;
}
.issue-flag-header {
  width: 4em;
}
th.check-header {
  width: 32px;
  label {
    display: block;
    width: 100%;
    height: 100%;
  }
}
td.check-column {
  text-align: center;
  padding: 0;
  label {
    display: block;
    width: 100%;
    height: 100%;
  }
}

.checkbox {
  width: 20px;
  height: 20px;
  margin: auto;
  vertical-align: text-bottom;
}
</style>
