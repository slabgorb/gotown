angular
  .module("gotown-town")
  .controller('TownController', [TownController, '$http'])

function TownController($http) {
    return {
      refresh: function() {
        $http( {
          method:'GET',
          url:'/town_names'
        }).then(function(response) {
          $scope.townNames = response.data;
        })
      }
    }
}
