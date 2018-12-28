import Vuex from 'vuex';
import { auth } from '@/services/firebase';

export const SESSION_LOGIN = 'session_login';
export const SESSION_LOGOUT = 'session_logout';

const store = new Vuex.Store({
  state: () => ({
    user: null,
    token: null
  }),
  mutations: {
    [SESSION_LOGIN](_, { user }) {
      this.user = user;
    },
    [SESSION_LOGOUT]() {
      this.user = null;
    }
  },
  actions: {}
});

auth.onAuthStateChanged(user => {
  console.log(user);
  if (user) {
    store.commit(SESSION_LOGIN, {
      user: {
        email: user.email
      }
    });
  } else {
    store.commit(SESSION_LOGOUT);
  }
});

export default () => store;
