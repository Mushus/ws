<template>
  <section>
    <b-breadcrumb>
      <b-breadcrumb-item :active="true">書類を印刷</b-breadcrumb-item>
    </b-breadcrumb>
    <b-navbar>
      <h2>書類印刷 - {{ date }}</h2>
      <b-button-toolbar>
        <b-button-group>
          <b-button :to="{ name: 'print-receipts-date', params: { date: prev } }">前の月へ</b-button>
          <b-button :to="{ name: 'print-receipts-date', params: { date: next } }">次の月へ</b-button>
        </b-button-group>
      </b-button-toolbar>
    </b-navbar>
    <div v-for="article in articles" :key="article.id">
      <b-container style="font-size: 1.2rem;">{{ article.data.name }}</b-container>
      <b-table
        :fixed="true"
        :items="getItems(article.id)"
        :fields="fields"
        >
        <template slot="tenantName" slot-scope="data">
          <span v-if="data.item.tenantName != null">
            {{ data.item.tenantName }}
          </span>
          <b-badge v-else>
            空室
          </b-badge>
        </template>
        <template slot="receipt" slot-scope="data">
          <b-badge v-if="data.item.receipt != null" variant="primary">
            印刷済み
          </b-badge>
          <b-badge v-else>
            未印刷
          </b-badge>
        </template>
        <template slot="control" slot-scope="data">
          <template v-if="data.item.tenantId != null">
            <b-button
              size="sm"
              variant="primary"
              :to="{
                name: 'print-receipts-date-tenantId',
                params: { date: date, tenantId: data.item.tenantId } }"
              >
              領収書を印刷
            </b-button>
          </template>
        </template>
      </b-table>
    </div>
  </section>
</template>

<script>
import Vue from "vue";
import { normalizeArticle, normalizeRoom, normalizeTenant, normalizeReceipt } from '@/util/normalize';

export default Vue.extend({
  async asyncData({ error, $firestore, params: { date } }) {
    // NOTE: numdate は YYYYMM の number 型
    const numDate = Number(date);
    const articlesRef = $firestore.collection('articles');
    const roomsRef = $firestore.collection('rooms');
    const tenantsRef = $firestore.collection('tenants');
    const receiptRef = $firestore.collection('receipts');
    const articles = [];
    const rooms = [];
    let tenants = [];
    let receipts = [];
    try {
      await Promise.all([
        (async () => {
          const articlesDoc = await articlesRef.get();
          articlesDoc.forEach(article =>
            articles.push(normalizeArticle(article.id, article.data())));
        })(),
        (async () => {
          const roomsDoc = await roomsRef.get();
          roomsDoc.forEach(room =>
            rooms.push(normalizeRoom(room.id, room.data())));
        })(),
        (async () => {
          // 指定月に入居中の部屋の一覧
          const tenantDoc = await tenantsRef
            // NOTE: 複数のフィールドに不等式は使用できない
            //.where('moveInAt', '<', numYearMonth * 100 + 99)
            .where('moveOutAt', '>', numDate * 100)
            .get();
          tenantDoc.forEach(tenant => {
            const tenantData = tenant.data();
            if (tenantData.moveInAt < numDate * 100 + 99) {
              tenants.push(normalizeTenant(tenant.id, tenantData));
            }
          });
          // 複数のフィールドに不等式が使えないので
          tenants = tenants.filter(tenant => tenant.data.moveInAt < numDate * 100 + 99)
        })(),
        (async () => {
          const receiptDoc = await receiptRef
            .where('date', '==', date)
            .get();
          receiptDoc.forEach(receipt => {
            receipts.push(normalizeReceipt(receipt.id, receipt.data()));
          });
        })()
      ]);
    } catch(e) {
      console.log(e);
      return error({ statusCode: 500, message: 'データ取得失敗' });
    }

    return {
      date: numDate,
      articles,
      rooms,
      tenants,
      receipts,
      fields: [
        {
          key: 'name',
          label: '部屋名',
        },
        {
          key: 'tenantName',
          label: '入居者',
        },
        {
          key: 'receipt',
          label: '領収書',
          class: 'w-25',
        },
        {
          key: 'control',
          label: '操作',
          class: 'w-25',
        },
      ],
    };
  },
  computed: {
    month() {
      return this.date % 100;
    },
    year() {
      return this.date / 100 | 0;
    },
    prev() {
      const month = this.month;
      const prevMonth = month === 1 ? 12 - 100 : month - 1;
      return this.year * 100 + prevMonth;
    },
    next() {
      const month = this.month;
      const nextMonth = month === 12 ? 1 + 100 : month + 1;
      return this.year * 100 + nextMonth;
    },
  },
  methods: {
    getItems(articleId) {
      // 部屋と入居者の Right join
      const rooms = [];
      this.rooms
        .filter(room => room.data.articleId === articleId)
        .forEach(room => {
          const tenants = this.tenants
            .filter(tenant => tenant.data.roomId === room.id)
          tenants.forEach(tenant => {
            const receipt = this.receipts.find(receipt => receipt.data.tenantId === tenant.id) || null;
            rooms.push({
              id: room.id,
              name: room.data.name,
              tenantName: tenant.data.name,
              tenantId: tenant.id,
              receipt,
            });
          });
          if (tenants.length === 0) {
            rooms.push({
              id: room.id,
              name: room.data.name,
              tenantName: null,
              tenantId: null,
              receipt: null,
            });
          }
        });
      return rooms;
    }
  }
})
</script>
