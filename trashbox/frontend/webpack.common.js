const HtmlWebpackPlugin = require('html-webpack-plugin');
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const path = require('path');

module.exports = {
  entry: [
    './src/index.tsx',
  ],
  module: {
    rules: [
      {
        test: /\.tsx?$/,
        use: [
          'ts-loader',
          {
            loader: 'eslint-loader',
            options: {
              fix: true,
            }
          }
        ],
        exclude: /node_modules/,
      },
      {
        test: /\.scss$/,
        use: [ MiniCssExtractPlugin.loader, 'css-loader', 'sass-loader' ],
      }
    ],
  },
  resolve: {
    extensions: [ '.tsx', '.ts', '.js', '.scss' ],
    alias: {
      '~': path.resolve(__dirname, 'src/'),
    },
  },
  output: {
    path: path.resolve('dist/'),
    filename: 'bundle.js',
  },
  plugins: [
    new MiniCssExtractPlugin({
      filename: '[name].css',
    }),
    new HtmlWebpackPlugin({
      title: 'React App',
      template: 'assets/index.html',
    }),
  ]
};
