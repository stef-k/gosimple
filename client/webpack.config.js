const path = require('path');

const PATHS = {
  name: 'bundle.js',
  app: path.join(__dirname, 'app'),
  // output the build template at Beego's views directory
  build: path.join('../views'),
  // output the build bundle.js file at Beego's static/js directory
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
    modulesDirectories: ["./app/components", "app/*/*", "node_modules"]
  },
  plugins: []
};
