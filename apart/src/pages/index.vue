class as section as template; = "container" >
<template>
  <section>
    <h2>物件管理</h2>
    <ul>
      <li>
        <nuxt-link :to="{ path: '/receipts', params: {} }">
          hoge
        </nuxt-link>
      </li>
    </ul>
  </section>
</template>

<script>
import Vue from 'vue';

export default Vue.extend({
})
</script>
