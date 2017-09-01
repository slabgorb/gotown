const path = require('path');

module.exports = {
  entry: {
    app: './web/src/app.jsx'
  },
  output: {
    filename: '[name].js',
    path: path.join(__dirname, './web/scripts'),
  },
  resolve: {
    modules: [path.resolve("web/src"), "node_modules"]

  },
  watch: true,
  module: {
    rules: [
      {
        test: /\.scss$/,
        exclude: /node_modules/,
        use: [
          {
            loader: "style-loader"
          },
          {
            loader: "css-loader", options: {sourceMap: true}
          },
          {
            loader: "sass-loader", options: {sourceMap: true}
          }
        ]
      },
      {
        test: /\.js$|\.jsx$/,
        exclude: /(node_modules|bower_components)/,
        use: {
          loader: 'babel-loader',
          options: {
            presets: ['react', 'es2015']
          }
        }
      }
    ]
  }
}