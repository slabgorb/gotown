const LiveReloadPlugin = require('webpack-livereload-plugin');

module.exports = {
  mode: 'development',
  entry: {
    app: '/web/src/app.jsx',
  },
  output: {
    filename: 'bundle.js',
    path: '/web/scripts',
  },
  resolve: {
    modules: ['/web/node_modules'],
    extensions: ['.js', '.jsx'],

  },
  resolveLoader: {
    modules: ['/web/node_modules'],
  },
  plugins: [
    new LiveReloadPlugin({}),
  ],
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
