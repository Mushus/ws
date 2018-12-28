import { auth, store } from '@/services/firebase';

export default (ctx, inject) => {
  // Inject firestore
  ctx.$firestore = store;
  inject('firestore', store);

  // Inject auth
  ctx.$auth = auth;
  inject('auth', auth);
};
