<template>
  <section>
    <b-container>
      <b-row class="justify-content-sm-center">
        <b-col sm="5">
          <h2>ログイン</h2>
          <div class="message-error" v-if="error !== ''">
            {{ error }}
          </div>
          <b-form @submit.prevent="login($event)">
            <b-form-group label="メールアドレス">
              <b-input type="text" v-model="email" required />
            </b-form-group>
            <b-form-group label="メールアドレス">
              <b-input type="password" v-model="password" required />
            </b-form-group>
            <b-button-group>
              <b-button variant="primary" type="submit" :disabled="isProcessing">ログイン</b-button>
            </b-button-group>
          </b-form>
        </b-col>
      </b-row>
    </b-container>
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
