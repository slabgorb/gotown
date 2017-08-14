module.exports = function(gulp, plugins) {
  return function() {
      plugins.go.run("../main.go", {cwd: __dirname, stdio:'inherit'})
  }
}
