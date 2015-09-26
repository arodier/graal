
angular.module("myapp", []).controller("MainController", function($scope, $http) {

    var hostnamePromise = $http.get("/net/hostname");

    hostnamePromise.success(function(data, status, headers, config) {
        // $scope.myData.fromServer = data.title;
        $scope.instance = {
            hostname: data.Data
        };
    });

    hostnamePromise.error(function(data, status, headers, config) {
        alert("AJAX failed!");
    });
});
