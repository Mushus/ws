import Vuex from 'vuex';
import { auth } from '@/services/firebase';

export const SESSION_LOGIN = 'session_login';
export const SESSION_LOGOUT = 'session_logout';

const store = new Vuex.Store({
  state: () => ({
    user: null,
    token: null
  }),
  getters: {
    loggedIn(state) {
      return state.user !== null;
    },
    user(state) {
      return {
        ...state.user
      };
    }
  },
  mutations: {
    [SESSION_LOGIN](state, { user }) {
      state.user = user;
    },
    [SESSION_LOGOUT]() {
      this.user = null;
    }
  },
  actions: {
    // ログイン
    [SESSION_LOGIN]({ commit }, { user }) {
      const routeName = this.app.context.route.name;
      commit(SESSION_LOGIN, { user });
      // ログイン画面からのログインではインデックスに飛ばす
      if (routeName === 'session-login') {
        this.$router.replace({ name: 'index' });
      }
    },
    // ログアウト
    [SESSION_LOGOUT]({ commit }) {
      const login = { name: 'session-login' };
      if (this.getters.loggedIn) {
        // ログインしている場合はログアウトを行う
        auth.signOut().then(() => {
          commit(SESSION_LOGOUT);
          this.$router.replace(login);
        });
      } else {
        this.$router.replace(login);
      }
    }
  }
});

auth.onAuthStateChanged(user => {
  if (user) {
    store.dispatch(SESSION_LOGIN, {
      user: {
        email: user.email
      }
    });
  } else {
    store.dispatch(SESSION_LOGOUT);
  }
});

export default () => store;
