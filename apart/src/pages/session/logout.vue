<template>
  <section>
    <button type="button" @click="e => logout()" :disabled="isProcessing">ログアウト</button>
  </section>
</template>

<script>
import Vue from 'vue';
import { mapMutations } from 'vuex';
import { SESSION_LOGOUT } from '@/store';

export default Vue.extend({
  data() {
    return {
      isProcessing: false,
    };
  },
  methods: {
    logout() {
      // ログアウトボタンが一度押されたら何度も押されないように無効化する
      this.isProcessing = true
      // ログアウトしたら鏡ページに遷移
      this.$store.dispatch(SESSION_LOGOUT, {
        callback: () => {
          this.$router.push({
            name: 'index',
          })
        }
      })
    }
  }
});
</script>
