<template>
  <div>
    <h2>ログイン</h2>
    <div>
      {{password}}
      <div><label>email: <input type="text" v-model="email"></label></div>
      <div><label>password: <input type="password" v-model="password"></label></div>
      <div><button type="button" @click="login()">login</button></div>
    </div>
  </div>
</template>

<script lang="ts">
import {
  Component,
  Vue,
  Prop
} from "nuxt-property-decorator"
import { State } from "vuex-class"

@Component({
  components: {}
})
export default class extends Vue {
  @Prop()
  email: string;

  @Prop()
  password: string;

  login() {
    this.$store.dispatch('login', {
      email: this.email,
      password: this.password,
    })
  }
}
</script>
<style scoped>
</style>
