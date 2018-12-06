const path = require('path');
const merge = require('webpack-merge');
const FaviconsWebpackPlugin = require('favicons-webpack-plugin');
const common = require('./webpack.common.js');

module.exports = merge(common, {
    mode: 'production',
    plugins: [
      new FaviconsWebpackPlugin({
        logo: path.resolve('assets/logo.png'),
        prefix: 'icons/',
        background: '#2196F3',
        title: 'React Redux TypeScript',
        icons: {
          android: true,
          appleIcon: true,
          appleStartup: true,
          coast: false,
          favicons: true,
          firefox: true,
          opengraph: false,
          twitter: true,
          yandex: false,
          windows: true
        }
      })
    ]
});
