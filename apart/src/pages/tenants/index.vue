<template>
  <section>
    <b-navbar class="pt-3">
      <h2>入居者一覧</h2>
      <b-button-toolbar>
        <b-button-group>
          <b-button :to="{ name: 'tenants-create' }" variant="primary">入居者を作成する</b-button>
        </b-button-group>
      </b-button-toolbar>
    </b-navbar>
    <b-table hover fixed="true" :fields="fields" :items="items">
      <template slot="control" slot-scope="data">
        <b-button @click="e => tenantLink(data.item.id)">編集</b-button>
      </template>
    </b-table>
  </section>
</template>

<script>
import Vue from "vue";

export default Vue.extend({
  async asyncData({ error, $firestore }) {
    const tenantsRef = $firestore.collection('tenants');
    const tenants = [];
    try {
      const tenantsDoc = await tenantsRef.get();
      tenantsDoc.forEach(tenant => tenants.push({
        id: tenant.id,
        data: tenant.data(),
      }));
    } catch(e) {
      console.log(e);
      return error({ statusCode: 500, message: 'データ取得失敗' });
    }
    return {
      fields: [
        {
          key: 'name',
          label: '名前'
        },
        {
          key: 'control',
          label: '操作',
          class: 'w-25'
        }
      ],
      tenants,
    };
  },
  computed: {
    items() {
      return this.tenants.map(article => ({
        id: article.id,
        ...article.data,
      }));
    },
  },
  methods: {
    tenantLink(id) {
      this.$router.push({ name: 'tenants-id', params: { id } });
      return false;
    }
  }
})
</script>

