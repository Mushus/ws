const path = require('path');
const PrettierPlugin = require('prettier-webpack-plugin');

module.exports = {
  mode: 'production',
  entry: path.resolve(__dirname, './src/index.js'),
  module: {
    rules: [
      {
        test: /\.html$/,
        loader: 'html-loader'
      },
      {
        test: /\.scss$/,
        use: [
          "css-loader",
          "sass-loader"
        ]
      }
    ]
  },
  plugins: [
    new PrettierPlugin()
  ],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src')
    }
  }
};
