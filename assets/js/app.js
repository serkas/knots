var app = angular.module('KnotsApp', ['ngSanitize',  'hc.marked']);
var postConfig = {headers: {'Content-Type': "application/json", 'Accept': "application/json"}};

app.controller('MainCtr', function($scope, $http) {
    initEmptyKnot($scope);
    $scope.knots = [];

    getKnots($http, $scope);

    $scope.postNew = function() {
        var knot = angular.copy($scope.new);
        knot.raw = editor.value();
        $http.post('/knots', JSON.stringify(knot), postConfig).then(function(){
            initEmptyKnot($scope);
            getKnots($http, $scope);
        });
    };

    $scope.remove = function(type, id) {
        $http.delete('/knots/'+id, {}).then(function(){
            getKnots($http, $scope);
        });
    };

});

function getKnots($http, $scope) {
    $http.get("/knots").then(function(response) {
        $scope.knots = response.data.knots;
    });
}

function initEmptyKnot($scope){
    $scope.new = {raw: '', title: ''};
}