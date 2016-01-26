var app = angular.module('KnotsApp', ['ngSanitize', 'hc.marked']);
var postConfig = {headers: {'Content-Type': "application/json", 'Accept': "application/json"}};

app.controller('MainCtr', function ($scope, $http) {
    initEmptyKnot($scope);
    $scope.knots = [];

    getKnots($http, $scope);

    $scope.postNew = function () {
        var knot = angular.copy($scope.new);
        knot.raw = angular.copy(editor.value());
        if (knot.id) {
            $http.put('/knots/' + knot.id, JSON.stringify(knot), postConfig).then(function () {
                $scope.empty();
                getKnots($http, $scope);
            });
        } else {
            $http.post('/knots', JSON.stringify(knot), postConfig).then(function () {
                $scope.empty();
                getKnots($http, $scope);
            });
        }

    };

    $scope.remove = function (type, knot) {
        $http.delete('/knots/' + knot.id, {}).then(function () {
            getKnots($http, $scope);
        });
    };

    $scope.edit = function (type, knot) {
        $scope.action = "update";
        $scope.new = knot;
        editor.value(knot.raw);
    };

    $scope.empty = function () {
        $scope.action = "new";
        initEmptyKnot($scope);
        editor.value("");
    }
});

function getKnots($http, $scope) {
    $http.get("/knots").then(function (response) {
        $scope.knots = response.data.knots;
    });
}

function initEmptyKnot($scope) {
    $scope.new = {raw: '', title: ''};
}