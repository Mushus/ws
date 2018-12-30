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
  data() {
    return {
      tenant: {
        id: null,
        data: {
          'name': '',
        }
      }
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
