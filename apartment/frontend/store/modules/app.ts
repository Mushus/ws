const state = {
  title: "",
  drawer: false,
};

const getters = {
  title: state => state.title,
  drawer: state => state.drawer,
};

const actions = {
  setTitle({ commit }, title: string) {
    commit('setTitle', title)
  },
  setDrawerState({ commit }, opened: boolean) {
    commit('setDrawerState', opened)
  },
};

const mutations = {
  setTitle(state, title: string) {
    state.title = title
  },
  setDrawerState(state, opened: boolean) {
    state.drawer = opened;
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}
