<template>
  <div>
    <header>
      <h1>物件管理システム</h1>
    </header>
    <Sidenav/>
    <main>
      <nuxt/>
    </main>
    <footer>

    </footer>
  </div>
</template>

<script lang="ts">
import {
  Component,
  Vue
} from "nuxt-property-decorator"
import Sidenav from "~/components/Sidenav"

@Component({
  components: {
    Sidenav
  }
})
export default class extends Vue {}
</script>
