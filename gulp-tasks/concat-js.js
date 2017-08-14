module.exports = function(gulp, plugins) {
  return function() {
    return gulp.src([ 'web/src/app.js', 'web/src/**/*module.js', 'web/src/**/*.js' ])
      .pipe(plugins.sourcemaps.init())
      .pipe(plugins.concat('app.js'))
      .pipe(plugins.sourcemaps.write())
      .pipe(gulp.dest('web/scripts'))
  }
}
