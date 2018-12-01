import * as firebase from 'firebase/app';
import 'firebase/auth';

const firebaseConfig = {
	apiKey: process.env.FIREBASE_API_KEY,
	authDomain: `${process.env.FIREBASE_PROJECT_ID}.firebaseapp.com`,
	projectId: process.env.FIREBASE_PROJECT_ID
};

if (!firebase.apps.length) {
	firebase.initializeApp(firebaseConfig);
}

export const state = () => ({
	loginUser: null,
	articles: []
});

export const mutations = {
	setArticles(state, articles) {
		state.articles = articles;
	},
	setLoginUser(state, loginUser) {
		state.loginUser = loginUser;
	}
};

export const actions = {
	async nuxtServerInit({ commit }, { app }) {
		/*const people = await app.$axios.$get(
	"./random-data.json"
	)*/
		commit("setArticles", [{ id: "hoge", name: "hogehoge" }]);
	},
	// ログイン
	async createUser({ commit }, { email: email, password: password }) {
		console.log(firebase)
		firebase.auth().createUserWithEmailAndPassword(email, password).catch(function(error) {
			// Handle Errors here.
			var errorCode = error.code;
			var errorMessage = error.message;
			// ...
		});
	}
};
