<template>
  <section>
    <b-breadcrumb>
      <b-breadcrumb-item :active="true">入居状況の確認</b-breadcrumb-item>
      <b-breadcrumb-item :to="{ name: 'occupants' }">建物を選択する</b-breadcrumb-item>
      <b-breadcrumb-item :active="true">部屋を選択する</b-breadcrumb-item>
    </b-breadcrumb>
    <b-navbar>
      <h2>入居状況の確認</h2>
    </b-navbar>
    <b-form-group label="建物名">
      <b-form-input type="text" :value="article.data.name" readonly />
    </b-form-group>
    <b-table
      :fields="fields"
      :items="items"
      >
      <template slot="status" slot-scope="data">
        <strong
          v-if="data.item.status === status.available"
          colo
          >
          <b-badge variant="success">入居可能</b-badge>
        </strong>
        <span v-if="data.item.status === status.unavailable && data.item.tenant != null">
          <template
            v-if="data.item.tenant.data.moveInAt >= minDate || data.item.tenant.data.moveOutAt < maxDate"
            >
            <span　v-if="data.item.tenant.data.moveInAt >= minDate">
              {{ numberToDate(data.item.tenant.data.moveInAt) }}
            </span>
            ～
            <span v-if="data.item.tenant.data.moveOutAt < maxDate">
              {{ numberToDate(data.item.tenant.data.moveOutAt) }}
            </span>
          </template>
          <template v-else>
            <b-badge variant="danger">入居中</b-badge>
          </template>
        </span>
      </template>
      <template slot="tenant" slot-scope="data">
        <span v-if="data.item.tenant != null">
          {{ data.item.tenant.data.name }}
        </span>
      </template>
      <template slot="control" slot-scope="data">
        <b-button-group size="sm">
          <b-button
            type="button"
            v-if="data.item.status === status.available"
            :to="{ name: 'occupants-articleId-roomId-movein', params: { articleId, roomId: data.item.id } }
            " variant="primary">入居する</b-button>
          <b-button
            type="button"
            v-if="data.item.status === status.unavailable && data.item.tenant != null"
            :to="{
              name: 'occupants-articleId-roomId-moveout-tenantId',
              params: { articleId, roomId: data.item.id, tenantId: data.item.tenant.id }
            }"
            variant="danger">退去する</b-button>
          <b-button
            type="button"
            :to="{ name: 'occupants-articleId-roomId-history', params: { articleId, roomId: data.item.id } }
            ">履歴を見る</b-button>
        </b-button-group>
      </template>
    </b-table>
  </section>
</template>

<script>
import Vue from 'vue';
import moment from 'moment';
import { MAX_DATE, DATE_FORMAT } from '@/util/constants';
import { normalizeArticle, normalizeRoom, normalizeTenant } from '@/util/normalize';

export default Vue.extend({
  async asyncData({ error, params: { articleId }, $firestore }) {
    const articlesRef = $firestore.collection('articles');
    const roomsRef = $firestore.collection('rooms');
    const tenantsRef = $firestore.collection('tenants');

    let article;
    const rooms = [];
    const tenants = [];
    try {
      const articleDoc = await articlesRef
        .doc(articleId)
        .get();

      if (!articleDoc.exists) {
        return error({ statusCode: 404, message: '指定された建物が存在しません' });
      }

      article = normalizeArticle(articleDoc.id, articleDoc.data());
      await Promise.all([
        (async () => {
          const roomsDoc = await roomsRef
            .where('articleId', '==', articleId)
            .orderBy('index')
            .get();
          roomsDoc.forEach(room =>
            rooms.push(normalizeRoom(room.id, room.data())));
        })(),
        (async () => {
          const tenantsDoc = await tenantsRef
            .where('articleId', '==', articleId)
            .where('moveOutAt', '>=', Number(moment().format(DATE_FORMAT)))
            .get();
          tenantsDoc.forEach(tenant =>
            tenants.push(normalizeTenant(tenant.id, tenant.data())));
        })()
      ]);
    } catch (e) {
      console.log(e);
      return error({ statusCode: 500, message: 'データ取得失敗' });
    }

    // 日付が今日より小さいデータはテーブルから省略
    const minDate = Number(moment().format(DATE_FORMAT));

    return {
      articleId,
      article,
      rooms,
      tenants,
      fields: [
        {
          key: 'name',
          label: '部屋名',
        },
        {
          key: 'status',
          label: '入居状況',
        },
        {
          key: 'tenant',
          label: '入居者',
        },
        {
          key: 'control',
          label: '操作',
        },
      ],
      status: {
        available: 0,
        unavailable: 1,
      },
      maxDate: MAX_DATE,
      minDate,
    };
  },
  computed: {
    items() {
      return this.rooms.map(room => {
        const tenant = this.tenants.find(tenant => tenant.data.roomId === room.id);
        const status = tenant? this.status.unavailable : this.status.available;
        return {
          id: room.id,
          name: room.data.name,
          status,
          tenant: tenant? tenant : null,
        };
      });
    },
  },
  methods: {
    numberToDate(num) {
      return num ? moment(String(num), DATE_FORMAT).format('YYYY年MM月DD日'): null;
    }
  },
})
</script>
