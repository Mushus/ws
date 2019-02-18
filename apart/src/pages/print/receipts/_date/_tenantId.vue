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
      <b-input-group append="円">
        <b-form-input
          type="number"
          v-model="receipt.data.rent"
          required
          min="0"
          />
      </b-input-group>
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
import { normalizeTenant, normalizeArticle, normalizeReceipt } from '@/util/normalize';
import { DATE_FORMAT } from '@/util/constants';

export default Vue.extend({
  components: {
    Datepicker,
  },
  async asyncData({ error, params: { tenantId, date }, $firestore }) {
    const receiptsRef = $firestore.collection('receipts');
    const tenantsRef = $firestore.collection('tenants');
    const articlesRef = $firestore.collection('articles');

    const result = {
      ja,
    };
    let receipt;
    try {
      const receipts = await receiptsRef
        .where('tenantId', '==', tenantId)
        .get();

      // 過去に印刷した領収書が存在してたらそれを復元する
      if (receipts.size > 0) {
        // 複数ある可能性はあるが基本的に1つ
        const receipt = receipts.docs[0];
        return {
          ...result,
          receipt: normalizeReceipt(receipt.id, receipt.data()),
        };
      }

      const tenantDoc = await tenantsRef
        .doc(tenantId)
        .get();

      if (!tenantDoc.exists) {
        return error({ statusCode: 404, message: '指定された入居者は存在しません' });
      }

      const tenant = normalizeTenant(tenantDoc.id, tenantDoc.data());

      const articleDoc = await articlesRef
        .doc(tenant.data.articleId)
        .get();

      if (!articleDoc.exists) {
        return error({ statusCode: 404, message: '指定された建物は存在しません' });
      }

      const article = normalizeArticle(articleDoc.id, articleDoc.data());

      return {
        ...result,
        receipt: normalizeReceipt(null, {
          publishAt: Number(moment().format(DATE_FORMAT)),
          tenantName: tenant.data.name,
          rent: tenant.data.rent,
          commonAreaCharge: tenant.data.commonAreaCharge,
          parkingFee: tenant.data.parkingFee,
          waterCharge: 0,
          administrator: article.data.administrator,
          tenantId: tenant.id,
          date,
        }),
      };
    } catch (e) {
      console.log(e);
      return error({ statusCode: 500, message: 'データ取得失敗' });
    }
  },
  mounted() {
    const win = this.$refs.printfield.contentWindow;
    win.document.write(html);
    this.updatePreview();
  },
  computed: {
    receiptPublishAt: {
      get() {
        const publishAt = this.receipt.data.publishAt;
        return publishAt ? moment(String(publishAt), DATE_FORMAT).toDate(): null;
      },
      set(v) {
        this.receipt.data.publishAt = v? Number(moment(v).format(DATE_FORMAT)) : null;
      }
    }
  },
  watch: {
    receipt: {
      handler() {
        this.updatePreview();
      },
      deep: true,
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
      const win = window.open();
      win.document.write(html);
      win.document.close();
      // NOTE: iframe を印刷しようとすると chrome が印刷ダイアログを表示してくれない
      setTimeout(() => {
        win.document.body.innerHTML = template(this.receipt.data);
        win.print();
      });

      this.submit();
    },
    async submit() {
      const receiptsRef = this.$firestore.collection('receipts');
      try {
        if (this.receipt.id) {
          const receiptId = this.receipt.id;
          const receipt = receiptsRef.doc(receiptId);
          await receipt.update(this.receipt.data);
        } else {
          const receipt = await receiptsRef.add(this.receipt.data);
        }
        const date = this.$route.params.date;
        this.$router.push({
          name: 'print-receipts-date',
          params: {
            date,
          },
        });
      } catch(e) {
        console.log(e);
      }
    }
  }
})
</script>
