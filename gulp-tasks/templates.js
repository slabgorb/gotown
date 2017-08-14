module.exports = function(gulp, plugins) {
  return function() {
    return gulp.src(['app/templates/**/*.html'])
      .pipe(plugins.templatecache('templates.js', {root:'app'}))
      .pipe(gulp.dest('app/scripts'))
  }
}
