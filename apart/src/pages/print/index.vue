<template>
</template>

<script>
import Vue from "vue";

export default Vue.extend({
  mounted() {
    this.$router.replace({ name: 'print-receipts-date', params: { date: '201901' } });
  },
})
</script>
