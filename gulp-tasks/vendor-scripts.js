var merge2 = require('merge2');
var bowerMain = require('bower-main');
var bowerMainJavaScriptFiles = bowerMain('js','min.js');

module.exports = function(gulp, plugins) {
  return function() {
    merge2(
      gulp.src(bowerMainJavaScriptFiles.minified),
      gulp.src(bowerMainJavaScriptFiles.minifiedNotFound)
        .pipe(plugins.concat('tmp.min.js'))
        .pipe(plugins.uglify()),
      gulp.src('bower_components/underscore.string/dist/underscore.string.min.js'),
      gulp.src('bower_components/angular-electron/d3.js')
    )
    .pipe(plugins.concat('vendor.min.js'))
    .pipe(gulp.dest('web/scripts'))
  }
}
