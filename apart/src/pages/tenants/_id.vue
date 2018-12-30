<template>
  <section>
    <b-navbar class="pt-3">
      <h2>入居者詳細</h2>
      <b-button-toolbar>
        <b-button-group>
          <b-button :to="{ name: 'tenants' }">一覧へ</b-button>
        </b-button-group>
      </b-button-toolbar>
    </b-navbar>
    <b-form @submit="e => (submit(), false)">
      <b-form-group label="氏名">
        <b-form-input type="text" v-model="tenant.data.name" required placeholder="山田太郎" />
      </b-form-group>
      <b-button-group>
        <b-button type="submit" variant="primary">作成する</b-button>
      </b-button-group>
    </b-form>
  </section>
</template>

<script>
import Vue from 'vue';

export default Vue.extend({
  async asyncData({ error, params, $firestore }) {
    const tenantId = params.id;
    const tenantRef = $firestore.collection('tenants').doc(tenantId);

    let tenant;
    try {
      const tenantDoc = await tenantRef.get();

      if (!tenantDoc.exists) {
        return error({ statusCode: 404, message: '指定された入居者が存在しません' });
      }

      tenant = {
        id: tenantId,
        data: tenantDoc.data(),
      };
    } catch(e) {
      console.log(e);
      return error({ statusCode: 500, message: 'データ取得失敗' });
    }
    return {
      tenant,
    };
  }
})
</script>
