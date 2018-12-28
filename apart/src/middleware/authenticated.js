export default ({ store, redirect }) => {
  if (!store.user) {
    redirect('session-login');
  }
};
