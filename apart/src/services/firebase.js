import * as firebase from 'firebase/app';
import 'firebase/firestore';
import 'firebase/auth';

const apiKey = process.env.FIREBASE_API_KEY;
const projectId = process.env.FIREBASE_PROJECT_ID;
const authDomain = `https://${projectId}.firebaseio.com`;

const config = {
  apiKey,
  authDomain,
  projectId
};

firebase.initializeApp(config);

const firestoreConfig = {
  timestampsInSnapshots: true
};
const firestore = firebase.firestore();
firestore.settings(firestoreConfig);

export const store = firestore;
export const auth = firebase.auth();
