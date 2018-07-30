const path = require('path');

module.exports = {
  mode: 'development',
  entry: {
    app: '/web/src/app.jsx',
  },
  output: {
    filename: 'bundle.js',
    path: '/app/scripts',
  },
  resolve: {
    modules: [path.resolve('/app'), 'node_modules'],
    extensions: ['.js', '.jsx'],

  },
  resolveLoader: {
    modules: [path.resolve('/app'), 'node_modules'],
  },
  watch: true,
  module: {
    rules: [
      {
        test: /\.scss$/,
        exclude: /node_modules/,
        use: [
          {
            loader: 'style-loader',
          },
          {
            loader: 'css-loader', options: { sourceMap: true },
          },
          {
            loader: 'sass-loader', options: { sourceMap: true },
          },
        ],
      },
      {
        test: /\.css$/,
        loader: 'style-loader!css-loader',
      },
      {
        test: /\.(woff|woff2|eot|ttf|otf)$/,
        use: [
          'file-loader',
        ],
      },
      {
        test: /\.js$|\.jsx$/,
        exclude: /node_modules/,
        use: {
          loader: 'babel-loader',
          options: {
            presets: ['react', 'es2015'],
          },
        },
      },
    ],
  },
};
