<template>
  <div class="app">
    <header class="app--header">
      <h1 class="app--header--title">Apartment</h1>
    </header>
    <nav class="app--nav">
      <ul class="app--nav--list">
        <li class="app--nav--list--link"><nuxt-link to="/">ダッシュボード</nuxt-link></li>
        <li class="app--nav--list--link"><nuxt-link to="/articles">物件の管理</nuxt-link></li>
        <li class="app--nav--list--link"><nuxt-link to="/tenants">入居者の管理</nuxt-link></li>
        <li class="app--nav--list--link"><nuxt-link to="/bills">請求書の発行</nuxt-link></li>
      </ul>
    </nav>
    <main>
      <nuxt/>
    </main>
    <footer>
      footer
    </footer>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator"
import { State, Action, Getter } from "vuex-class"

@Component({})
export default class extends Vue {
  @Getter('drawer', { namespace: 'app' }) drawerGetter
  @Getter('title', { namespace: 'app' }) title
  @Action('setDrawerState', { namespace: 'app' }) setDrawerState

  get drawer() {
    return this.drawerGetter
  }
  set drawer(val: boolean) {
    this.setDrawerState(val)
  }

  head() {
    return {
      titleTemplate: `${this.title}`,
    }
  }
}

</script>
