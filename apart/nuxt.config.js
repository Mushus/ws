const pkg = require('./package');
const extendConfig = require('./webpack.config.extend');

module.exports = {
  mode: 'spa',

  srcDir: './src',
  /*
  ** Headers of the page
  */
  head: {
    title: pkg.name,
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: pkg.description }
    ],
    link: [{ rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }]
  },

  /*
  ** Customize the progress-bar color
  */
  loading: { color: '#fff' },

  /*
  ** Global CSS
  */
  css: ['@/assets/css/main.scss'],

  /*
  ** Plugins to load before mounting the App
  */
  plugins: ['@/plugins/firebase'],

  /*
  ** Nuxt.js modules
  */
  modules: [
    'bootstrap-vue/nuxt',
    '@nuxtjs/axios',
    ['@nuxtjs/dotenv', { path: __dirname }]
  ],

  /*
  ** Build configuration
  */
  build: {
    /*
    ** You can extend webpack config here
    */
    extend(config, ctx) {
      extendConfig(config, ctx);
    }
  },
  router: {
    mode: 'hash'
  }
};
