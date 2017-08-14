var options = {
  fontsDir: 'fonts',
  cssDir: 'styles/scss',
  cssFilename: 'fonts.scss'

}
module.exports = function(gulp, plugins) {
  return function() {
    return gulp.src("gulp-tasks/fonts.list").
      pipe(plugins.googleWebfonts())
      pipe(gulp.dest('web'))

  }
}
