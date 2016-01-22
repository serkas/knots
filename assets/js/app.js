var app = angular.module('KnotsApp', []);
var postConfig = {headers: {'Content-Type': "application/json", 'Accept': "application/json"}};

app.controller('MainCtr', function($scope, $http) {
    initEmptyKnot($scope);
    $scope.knots = [];

    getKnots($http, $scope);

    $scope.postNew = function() {
        $http.post('/knots', JSON.stringify($scope.new), postConfig).then(function(){
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
    $scope.new = {text: '', title: ''};
}