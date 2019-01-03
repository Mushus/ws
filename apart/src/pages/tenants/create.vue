<template>
  <section>
    <b-navbar class="pt-3">
      <h2>入居者作成</h2>
      <b-button-toolbar>
        <b-button-group>
          <b-button :to="{ name: 'tenants' }">一覧へ</b-button>
        </b-button-group>
      </b-button-toolbar>
    </b-navbar>
    <b-form @submit="e => submit()">
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
    const articlesRef = $firestore.collection('articles');
    const roomsRef = $firestore.collection('article');

    let tenant;
    let articles = [];
    let rooms = [];
    try {
      const tenantDoc = await tenantRef.get();
      // 記事が存在しなければ表示できないので404にする
      if (!tenantDoc.exists) {
        return error({ statusCode: 404, message: '指定された入居者が存在しません' });
      }

      tenant = {
        id: tenantDoc.id,
        data: tenantDoc.data(),
      };

      await Promise.all([
        (async () => {
          const articlesDoc = await articlesRef.get();
          articlesDoc.forEach(article => articles.push({
            id: article.id,
            data: article.data(),
          }));
        }),
        (async () => {
          const roomsDoc = await roomsRef.get();
          roomsDoc.forEach(room => rooms.push({
            id: room.id,
            data: room.data(),
          }));
        })
      ]);
    } catch(e) {
      return error({ statusCode: 500, message: 'データ取得失敗' });
    }
    return {
      tenant,
      articles,
      rooms,
    };
  },
  methods: {
    async submit() {
      const tenantRef = this.$firestore.collection('tenants');
      try {
        const tenant = await tenantRef.add(this.tenant.data);
        this.$router.push({
          name: 'tenants-id',
          params: {
            id: tenant.id,
          },
        });
      } catch (e) {
        console.log(e);
      }
    }
  }
})
</script>
