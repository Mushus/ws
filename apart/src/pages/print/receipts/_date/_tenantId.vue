<template>
  <section>
    <b-breadcrumb>
      <b-breadcrumb-item :to="{ name: 'print' }">書類を印刷</b-breadcrumb-item>
      <b-breadcrumb-item :active="true">領収書を印刷</b-breadcrumb-item>
    </b-breadcrumb>
    <b-navbar>
      <h2>領収書を印刷</h2>
      <b-button-toolbar>
        <b-button-group>
          <b-button @click="print()" variant="primary">印刷する</b-button>
        </b-button-group>
      </b-button-toolbar>
    </b-navbar>
    <b-form-group label="発行日">
      <datepicker
        format="yyyy年MM月dd日"
        :language="ja"
        :bootstrap-styling="true"
        v-model="receiptPublishAt" />
    </b-form-group>
    <b-form-group label="入居者名">
      <b-form-input
        type="text"
        v-model="receipt.data.tenantName"
        required
        />
    </b-form-group>
    <b-form-group label="家賃">
      <b-form-input
        type="number"
        v-model="receipt.data.rent"
        required
        min="0"
        />
    </b-form-group>
    <b-form-group label="共益費">
      <b-input-group append="円">
        <b-form-input
          type="number"
          v-model="receipt.data.commonAreaCharge"
          required
          min="0"
          />
      </b-input-group>
    </b-form-group>
    <b-form-group label="駐車場料金">
      <b-input-group append="円">
        <b-form-input
          type="number"
          v-model="receipt.data.parkingFee"
          required
          min="0"
          />
      </b-input-group>
    </b-form-group>
    <b-form-group label="水道料金">
      <b-input-group append="円">
        <b-form-input
          type="number"
          v-model="receipt.data.waterCharge"
          required
          min="0"
          />
      </b-input-group>
    </b-form-group>
    <b-form-group label="管理者">
      <b-form-input
        type="text"
        v-model="receipt.data.administrator"
        required
        />
    </b-form-group>
    <iframe ref="printfield" frameborder="0" style="width: 100%; height: 500px;"></iframe>
  </section>
</template>

<script>
import Vue from 'vue';
import { ja } from 'vuejs-datepicker/dist/locale';
import Datepicker from 'vuejs-datepicker';
import moment from 'moment';
import { html, template } from '@/util/print/receipt';
const DateFormatString = 'YYYYMMDD';

export default Vue.extend({
  components: {
    Datepicker,
  },
  async asyncData({ error, params: { date } }) {
    return {
      receipt: {
        id: null,
        data: {
          publishAt: 20180101,
          tenantName: '入居者名',
          rent: 10000,
          commonAreaCharge: 3000,
          parkingFee: 1000,
          waterCharge: 1500,
          administrator: '管理者名',
        },
      },
      ja,
    }
  },
  mounted() {
    this.$refs.printfield.contentWindow.document.write(html);
    this.updatePreview();
  },
  computed: {
    receiptPublishAt: {
      get() {
        const publishAt = this.receipt.data.publishAt;
        return publishAt ? moment(String(publishAt), DateFormatString).toDate(): null;
      },
      set(v) {
        this.receipt.data.publishAt = v? Number(moment(v).format(DateFormatString)) : null;
      }
    }
  },
  methods: {
    updatePreview() {
      const iframeElement = this.$refs.printfield;
      const iframe = iframeElement.contentWindow;
      iframe.document.body.innerHTML = template(this.receipt.data);
      // iframe をスクロールバーに表示しない
      const iframeHtmlElem = iframe.document.documentElement;
      iframeElement.style.width = `${Math.ceil(iframeHtmlElem.scrollWidth)}px`;
      iframeElement.style.height = `${Math.ceil(iframeHtmlElem.scrollHeight)}px`;
    },
    print() {
      this.$refs.printfield.contentWindow.print();
    }
  }
})
</script>
