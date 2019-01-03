<template>
  <section>
    <h2>ログイン</h2>
    <section>
      <div class="message-error" v-if="error !== ''">
        {{ error }}
      </div>
      <form @submit.prevent="login($event)">
        <dl>
          <dt><label for="email">メールアドレス</label></dt>
          <dd><input type="email" id="email" v-model="email"></dd>
          <dt><label for="password">パスワード</label></dt>
          <dd><input type="password" id="password" v-model="password"></dd>
        </dl>
        <button type="submit" :disabled="isProcessing">ログイン</button>
      </form>
    </section>
  </section>
</template>

<script>
import Vue from "vue";

export default Vue.extend({
  layout: 'anonymous',
  data() {
    return {
      isProcessing: false,
      email: '',
      password: '',
      error: '',
    }
  },
  methods: {
    login(event) {
      if (event != null) event.preventDefault();
      // 通信中にエラーを消えるようにする
      this.error = '';
      this.isProcessing = true;

      const email = this.email;
      const password = this.password;
      this.$auth.signInWithEmailAndPassword(email, password).catch(e => {
        const errorCode = e.code;
        const errorMessage = e.message;
        // 画面にエラーを表示する
        this.error = `${errorCode}: ${errorMessage}`;
        this.isProcessing = false;
      });

      return false;
    }
  }
});
</script>
