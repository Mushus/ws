<template>
  <section>
    <b-breadcrumb>
      <b-breadcrumb-item :active="true">入居状況の確認</b-breadcrumb-item>
      <b-breadcrumb-item :to="{ name: 'occupants' }">建物を選択する</b-breadcrumb-item>
      <b-breadcrumb-item :to="{ name: 'occupants-articleId', params: { articleId } }">部屋を選択する</b-breadcrumb-item>
      <b-breadcrumb-item :active="true">入居する</b-breadcrumb-item>
    </b-breadcrumb>
    <b-navbar>
      <h2>入居する</h2>
    </b-navbar>
    <b-container>
      <p>入居する人を設定してください</p>
    </b-container>
    <b-form-group label="建物名">
      <b-form-input type="text" :value="article.data.name" readonly />
    </b-form-group>
    <b-form-group label="部屋名">
      <b-form-input type="text" :value="room.data.name" readonly />
    </b-form-group>
    <b-form @submit.prevent="submit()" class="pt-5">
      <b-form-group label="入居者名">
        <b-form-input type="text" v-model="tenant.data.name" required />
      </b-form-group>
      <b-form-group label="家賃">
        <b-input-group append="円">
          <b-form-input
            type="number"
            v-model="tenant.data.rent"
            required
            min="0"
            />
        </b-input-group>
      </b-form-group>
      <b-form-group label="共益費">
        <b-input-group append="円">
          <b-form-input
            type="number"
            v-model="tenant.data.commonAreaCharge"
            required
            min="0"
            />
        </b-input-group>
      </b-form-group>
      <b-form-group label="駐車料">
        <b-input-group append="円">
          <b-form-input
            type="number"
            v-model="tenant.data.parkingFee"
            required
            min="0"
            />
        </b-input-group>
      </b-form-group>
      <b-form-group label="入居日">
        <datepicker
          format="yyyy年MM月dd日"
          :language="ja"
          :bootstrap-styling="true"
          v-model="tenantMoveInAt" />
      </b-form-group>
      <b-button-group class="pt-3">
        <b-button type="submit" variant="primary">入居する</b-button>
      </b-button-group>
    </b-form>
  </section>
</template>

<script>
import Vue from 'vue';
import moment from 'moment';
import Datepicker from 'vuejs-datepicker';
import { ja } from 'vuejs-datepicker/dist/locale';
import { MAX_DATE, DATE_FORMAT } from '@/util/constants';
import { normalizeArticle, normalizeRoom, normalizeTenant } from '@/util/normalize';

export default Vue.extend({
  components: {
    Datepicker,
  },
  async asyncData({ error, $firestore, params: { articleId, roomId } }) {
    const articlesRef = $firestore.collection('articles');
    const roomsRef = $firestore.collection('rooms');

    let article;
    let room;
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
    } catch(e) {
      console.log(e);
      return error({ statusCode: 500, message: 'データ取得失敗' });
    }

    return {
      articleId,
      roomId,
      article,
      room,
      tenant: normalizeTenant(null, {
        articleId,
        roomId,
        name: '',
        rent: room.data.rent,
        commonAreaCharge: article.data.commonAreaCharge,
        parkingFee: article.data.parkingFee,
        moveInAt: Number(moment().format(DATE_FORMAT)),
        moveOutAt: MAX_DATE,
      }),
      ja,
    };
  },
  computed: {
    tenantMoveInAt: {
      get() {
        const moveInAt = this.tenant.data.moveInAt;
        return moveInAt ? moment(String(moveInAt), DATE_FORMAT).toDate(): null;
      },
      set(v) {
        this.tenant.data.moveInAt = v? Number(moment(v).format(DATE_FORMAT)) : null;
      }
    }
  },
  methods: {
    async submit() {
      const tenantsRef = this.$firestore.collection('tenants');
      await tenantsRef.add(this.tenant.data);
      this.$router.push({ name: 'occupants-articleId', params: { articleId: this.articleId } });
    }
  }
});
</script>
