<template>
  <v-layout align-center justify-center>
    <v-flex xs12 sm6 md6>
      <v-card class="elevation-12">
        <v-toolbar dark color="primary">
          <v-toolbar-title>ログイン</v-toolbar-title>
          <v-spacer></v-spacer>
        </v-toolbar>
        <v-card-text>
          <v-form>
            <v-text-field prepend-icon="person" label="ログイン名" type="text"></v-text-field>
            <v-text-field id="password" prepend-icon="lock" label="パスワード" type="password"></v-text-field>
          </v-form>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="primary">Login</v-btn>
        </v-card-actions>
      </v-card>
    </v-flex>
  </v-layout>
</template>

<script lang="ts">
import {
  Component,
  Vue,
  Prop
} from "nuxt-property-decorator"
@Component({})
export default class extends Vue {
  layout() {
    return 'anonymous'
  }
}
</script>

<style scoped>
</style>
