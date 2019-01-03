<template>
  <div v-if="loggedIn">
    <b-navbar toggleable="md" type="dark" variant="info">
      <b-navbar-brand :to="{ name: 'index'}">物件管理システム</b-navbar-brand>
      <b-navbar-nav class="ml-auto">
        <b-nav-item-dropdown :text="user.email" right>
          <b-dropdown-item @click="$router.push({ name: 'session-logout' })">ログアウト</b-dropdown-item>
        </b-nav-item-dropdown>
      </b-navbar-nav>
    </b-navbar>
    <b-container fluid class="pb-5">
      <b-row>
        <b-col class="pl-0 pr-0" cols="2">
          <b-nav vertical class="pt-2 sidenav">
            <b-nav-item :to="{ name: 'index' }">ダッシュボード</b-nav-item>
            <b-nav-item :to="{ name: 'articles' }">物件管理</b-nav-item>
            <b-nav-item :to="{ name: 'tenants' }">入居者管理</b-nav-item>
            <b-nav-item :to="{ name: 'occupants' }">入居状況の確認</b-nav-item>
            <b-nav-item :to="{ name: 'print' }">書類印刷</b-nav-item>
          </b-nav>
        </b-col>
        <b-col>
          <b-container tag="main">
            <nuxt />
          </b-container>
        </b-col>
      </b-row>
    </b-container>
  </div>
</template>

<script>
import Vue from 'vue';

export default Vue.extend({
  computed: {
    user() {
      return this.$store.getters.user;
    },
    loggedIn() {
      return this.$store.getters.loggedIn;
    }
  }
})
</script>
