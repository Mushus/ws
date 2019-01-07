<template>
  <section>
    <b-breadcrumb>
      <b-breadcrumb-item :to="{ name: 'tenants' }">入居者一覧</b-breadcrumb-item>
      <b-breadcrumb-item :active="true">入居者詳細</b-breadcrumb-item>
    </b-breadcrumb>
    <b-navbar>
      <h2>入居者詳細</h2>
    </b-navbar>
    <b-form @submit.prevent="submit()">
      <b-form-group label="氏名">
        <b-form-input type="text" v-model="tenant.data.name" required placeholder="山田太郎" />
      </b-form-group>
      <b-button-group>
        <b-button type="submit" variant="primary">更新する</b-button>
      </b-button-group>
    </b-form>
  </section>
</template>

<script>
import Vue from 'vue';
import { normalizeTenant } from '@/util/normalize';

export default Vue.extend({
  async asyncData({ error, params: { tenantId }, $firestore }) {
    const tenantRef = $firestore.collection('tenants').doc(tenantId);

    let tenant;
    try {
      const tenantDoc = await tenantRef.get();

      if (!tenantDoc.exists) {
        return error({ statusCode: 404, message: '指定された入居者が存在しません' });
      }

      tenant = normalizeTenant(tenantId, tenantDoc.data());
    } catch(e) {
      console.log(e);
      return error({ statusCode: 500, message: 'データ取得失敗' });
    }
    return {
      tenant,
    };
  },
  methods: {
    submit() {
      const tenantsRef = this.$firestore.collection('tenants');
      const tenantDoc = tenantsRef.doc(this.tenant.id);

      try {
        tenantDoc.update(this.tenant.data);
      } catch(e) {
        return console.log(e);
      }

      return false;
    }
  }
})
</script>
