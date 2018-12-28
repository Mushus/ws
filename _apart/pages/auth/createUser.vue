<template>
  <div>
    <h2>ユーザーを作成する</h2>
    <div>
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
  email: string = "";
  password: string = "";

  login() {
    console.log(this.email);
    this.$store.dispatch('createUser', {
      email: this.email,
      password: this.password,
    })
  }
}
</script>
<style scoped>
</style>
