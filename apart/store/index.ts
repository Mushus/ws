export const state = () => ({
	articles: []
})

export const mutations = {
  setArticles(state, articles) {
    state.articles = articles
  }
}

export const actions = {
  async nuxtServerInit({ commit }, { app }) {
    /*const people = await app.$axios.$get(
      "./random-data.json"
    )*/
    commit("setArticles", [
		{id:"hoge", name:"hogehoge"}
	])
  }
}
