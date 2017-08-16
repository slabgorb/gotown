module.exports = function(gulp, plugins) {
  return function() {
    return gulp.src(['web/templates/**/*.html'])
      .pipe(plugins.templatecache('templates.js', {root:'app'}))
      .pipe(gulp.dest('web/scripts'))
  }
}
