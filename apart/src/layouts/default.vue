<template>
  <div class="container" v-if="loggedIn">
    <header>
      <h1>物件管理システム</h1>
    </header>
    <nav>
      <ul>
        <li><nuxt-link :to="{ name: 'articles' }">物件管理</nuxt-link></li>
        <li><nuxt-link :to="{ name: 'receipts' }">領収書印刷</nuxt-link></li>
        <li><nuxt-link :to="{ name: 'session-logout' }">ログアウト</nuxt-link></li>
      </ul>
    </nav>
    <main>
      <nuxt />
    </main>
    <footer />
  </div>
</template>

<script>
import Vue from 'vue';

export default Vue.extend({
  computed: {
    loggedIn() {
      return this.$store.getters.loggedIn
    }
  }
})
</script>
