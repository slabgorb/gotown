var gulp = require('gulp');

var plugins = require('gulp-load-plugins')({
  rename: {
    'gulp-angular-templatecache': 'templatecache',
    'gulp-concat-css': 'concatCss'
  }
});


function getTask(task) {
    return require('./gulp-tasks/' + task)(gulp, plugins);
}

gulp.task('styles', getTask('styles'));
gulp.task('concatJS', getTask('concat-js'));
gulp.task('templates', getTask('templates'));
gulp.task('vendorScripts', getTask('vendor-scripts'));
gulp.task('webFonts', getTask('web-fonts'));
gulp.task('run', getTask('run'));

gulp.task('watch', function() {
  gulp.watch('web/styles/**/*.scss',['styles']);
  gulp.watch('web/src/**/*.js', ['concatJS']);
  gulp.watch('web/templates/**/*.html', ['templates']);
})


gulp.task('default', ['styles', 'concatJS', 'templates','watch', 'vendorScripts', 'webFonts']);
