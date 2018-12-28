<template>
  <section>
    <h2>ログイン</h2>
    <section>
      <div class="message-error" v-if="error !== ''">
        {{ error }}
      </div>
      <form @submit="e => (login(), false)">
        <dl>
          <dt><label for="email">メールアドレス</label></dt>
          <dd><input type="email" id="email" v-model="email"></dd>
          <dt><label for="password">パスワード</label></dt>
          <dd><input type="password" id="password" v-model="password"></dd>
        </dl>
        <button type="submit">ログイン</button>
      </form>
    </section>
  </section>
</template>

<script>

import Vue from "vue";

export default Vue.extend({
  data() {
    return {
      email: '',
      password: '',
      error: '',
    }
  },
  methods: {
    login() {
      // 通信中にエラーを消えるようにする
      this.error = '';
      const email = this.email;
      const password = this.password;
      this.$auth.signInWithEmailAndPassword(email, password).catch(e => {
        const errorCode = e.code;
        const errorMessage = e.message;
        // 画面にエラーを表示する
        this.error = `${errorCode}: ${errorMessage}`;
      });
    }
  },
  layout: 'anonymous'
});
</script>
