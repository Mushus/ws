<template>
  <section>
    <b-breadcrumb>
      <b-breadcrumb-item :active="true">入居状況の確認</b-breadcrumb-item>
      <b-breadcrumb-item :to="{ name: 'occupants' }">建物を選択する</b-breadcrumb-item>
      <b-breadcrumb-item :to="{ name: 'occupants-articleId', params: { articleId } }">部屋を選択する</b-breadcrumb-item>
      <b-breadcrumb-item :active="true">退去する</b-breadcrumb-item>
    </b-breadcrumb>
    <b-navbar>
      <h2>退去する</h2>
    </b-navbar>
    <b-container>
      <p>退去日時を設定してください</p>
    </b-container>
    <b-form-group label="建物名">
      <b-form-input type="text" :value="article.data.name" readonly />
    </b-form-group>
    <b-form-group label="部屋名">
      <b-form-input type="text" :value="room.data.name" readonly />
    </b-form-group>
    <b-form-group label="入居者名">
      <b-form-input type="text" v-model="tenant.data.name" readonly />
    </b-form-group>
    <b-form-group label="入居日">
      <b-form-input type="text" v-model="tenantMoveInAt" readonly />
    </b-form-group>
    <b-form @submit.prevent="submit()" class="pt-5">
      <b-form-group label="退去日">
        <datepicker
          format="yyyy年MM月dd日"
          :language="ja"
          :bootstrap-styling="true"
          v-model="tenantMoveOutAt"
          required />
      </b-form-group>
      <b-button-group class="pt-3">
        <b-button type="submit" variant="primary">退去する</b-button>
      </b-button-group>
    </b-form>
  </section>
</template>

<script>
import Vue from 'vue';
import moment from 'moment';
import Datepicker from 'vuejs-datepicker';
import { ja } from 'vuejs-datepicker/dist/locale';
import { DATE_FORMAT } from '@/util/constants';
import { normalizeArticle, normalizeRoom, normalizeTenant } from '@/util/normalize';

export default Vue.extend({
  components: {
    Datepicker,
  },
  async asyncData({ error, $firestore, params: { articleId, roomId, tenantId } }) {
    const articlesRef = $firestore.collection('articles');
    const roomsRef = $firestore.collection('rooms');
    const tenantsRef = $firestore.collection('tenants');

    let article;
    let room;
    let tenant;
    try {
      const articleDoc = await articlesRef.doc(articleId).get();
      if (!articleDoc.exists) {
        return error({ statusCode: 404, message: '指定された建物は存在しません' });
      }
      article = normalizeArticle(articleDoc.id, articleDoc.data());

      const roomDoc = await roomsRef.doc(roomId).get();
      if (!roomDoc.exists) {
        return error({ statusCode: 404, message: '指定された部屋は存在しません' });
      }
      room = normalizeRoom(roomDoc.id, roomDoc.data());

      const tenantDoc = await tenantsRef.doc(tenantId).get();
      if (!tenantDoc.exists) {
        return error({ statusCode: 404, message: '指定された入居者は存在しません' });
      }
      tenant = normalizeTenant(tenantDoc.id, {
        ...tenantDoc.data(),
        moveOutAt: Number(moment().format(DATE_FORMAT)),
      });
    } catch(e) {
      console.log(e);
      return error({ statusCode: 500, message: 'データ取得失敗' });
    }

    return {
      articleId,
      roomId,
      article,
      room,
      tenant,
      ja,
    };
  },
  computed: {
    tenantMoveInAt: {
      get() {
        return moment(this.tenant.data.moveInAt).format('YYYY年MM月DD');
      },
      set() {},
    },
    tenantMoveOutAt: {
      get() {
        const moveOutAt = this.tenant.data.moveOutAt;
        return moveOutAt ? moment(String(moveOutAt), DATE_FORMAT).toDate(): null;
      },
      set(v) {
        this.tenant.data.moveOutAt = v? Number(moment(v).format(DATE_FORMAT)) : null;
      }
    }
  },
  methods: {
    async submit() {
      const tenantId = this.tenant.id;
      const tenantsRef = this.$firestore.collection('tenants');
      await tenantsRef.doc(tenantId).update(this.tenant.data);
      this.$router.push({ name: 'occupants-articleId', params: { articleId: this.articleId } });
    }
  }
});
</script>
