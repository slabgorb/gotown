var options = {
  fontsDir: 'styles/fonts',
  cssDir: 'styles',
  cssFilename: 'fonts.scss'

}
module.exports = function(gulp, plugins) {

  return function() {
    return gulp.src("./fonts.list").
      pipe(plugins.googleWebfonts(options))

  }
}
