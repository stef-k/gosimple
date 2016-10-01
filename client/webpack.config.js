const path = require('path');

const PATHS = {
  name: 'bundle.js',
  app: path.join(__dirname, 'app'),
  build: path.join('../views'),
  out: path.join('../static', 'js')
};

module.exports = {
  entry: [
    './app/index.jsx'
  ],

  output: {
    path: PATHS.out,
    filename: PATHS.name
  },

  module: {
    loaders: [{
      test: /\.jsx?$/,
      exclude: /node_modules/,
      loader: 'babel-loader',
      query: {
        presets: ['es2015', 'react']
      }
    }]
  },
  resolve: {
    extensions: ["", ".jsx", ".js"],
    modulesDirectories: ["./app/components", "node_modules"]
  },
  plugins: []
};
