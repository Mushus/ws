const webpack = require('webpack');
const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin')

module.exports = {
	module: {
		rules: [
            {
                test: /\.tsx?$/,
                use: 'ts-loader',
                exclude: /node_modules/
            },
			{
				test: /\.(scss|css)$/,
				use: [
					{ loader: 'style-loader' },
					{ loader: 'css-loader' },
					{ loader: 'sass-loader' }
				]
			},
			{
				test: /\.html$/,
				use: "html-loader"
			}
		]
	},

	entry: {
		app: './src/index.tsx',
    },
    resolve: {
        extensions: [ '.tsx', '.ts', '.js' ]
    },

	output: {
		filename: '[name].js',
		path: path.resolve(__dirname, 'dist')
  },

	mode: 'development',
	optimization: {
		splitChunks: {
			cacheGroups: {
				vendors: {
					priority: -10,
					test: /[\\/]node_modules[\\/]/
				}
			},
			chunks: 'async',
			minChunks: 1,
			minSize: 30000,
			name: true
		}
	},
	plugins: [
		new HtmlWebpackPlugin({
		    template: "./html/index.html"
		})
	]
};
