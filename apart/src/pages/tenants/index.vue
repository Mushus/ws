<template>
  <section>
    <b-breadcrumb>
      <b-breadcrumb-item :active="true">入居者一覧</b-breadcrumb-item>
    </b-breadcrumb>
    <b-navbar>
      <h2>入居者一覧</h2>
    </b-navbar>
    <b-table :fixed="true" :fields="fields" :items="items">
      <template slot="control" slot-scope="data">
        <b-button :to="{ name: 'tenants-tenantId', params: { id: data.item.id } }">編集する</b-button>
      </template>
    </b-table>
  </section>
</template>

<script>
import Vue from "vue";
import { normalizeArticle, normalizeRoom, normalizeTenant, normalizeReceipt } from '@/util/normalize';

export default Vue.extend({
  async asyncData({ error, $firestore }) {
    const tenantsRef = $firestore.collection('tenants');
    const tenants = [];
    try {
      const tenantsDoc = await tenantsRef.get();
      tenantsDoc.forEach(tenant =>
        tenants.push(normalizeTenant(tenant.id, tenant.data())));
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

