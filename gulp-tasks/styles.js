module.exports = function(gulp, plugins) {
  return function() {
    return gulp.src('web/styles/scss/**/*.scss')
      .pipe(plugins.sass().on('error', plugins.sass.logError))
      .pipe(gulp.dest('web/styles/'))
      .pipe(plugins.concatCss("styles/bundle.css"))
      .pipe(gulp.dest('web'));
  }
}
