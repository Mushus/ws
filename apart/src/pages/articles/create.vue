<template>
  <section>
    <h2>物件作成</h2>
    <form @submit="submit">
      <section>
        <label for="name">物件名</label>
        <input id="name" type="text" v-model="input.name" />
      </section>
      <section>
        <button type="submit">作成する</button>
      </section>
    </form>
  </section>
</template>

<script>
import Vue from "vue"

export default Vue.extend({
  data() {
    return {
      input: {
        name: ''
      }
    }
  },
  methods: {
    submit() {
      this.$firestore
        .collection('articles')
        .add(this.input)
        .then(e => console.log(e))
        .catch(e => console.log(e))
      return false
    }
  }
})
</script>
