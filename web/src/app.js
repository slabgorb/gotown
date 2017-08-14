
angular
.module("gotown", ["gotown-town"])
.controller('TownController', function($http, $scope) {
  var ctrl = this;
  ctrl.refresh = function() {
    $http( {
      method:'GET',
      url:'/town_names'
    }).then(function(response) {
      $scope.townNames = _.uniq(response.data);
    })

  }
  ctrl.refresh()
})
