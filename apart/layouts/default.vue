<template>
  <div>
    <header>

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
import Sidenav from "~/components/Sidenav.vue"

@Component({
  components: {
    Sidenav
  }
})
export default class extends Vue {}
</script>
