var app = angular.module('KnotsApp', []);
var postConfig = {headers: {'Content-Type': "application/json", 'Accept': "application/json"}};

app.controller('MainCtr', function($scope, $http) {
    $scope.desc = 'Knots of wisdom';

    $scope.new = initNewPost();

    $scope.postNew = function() {
        $http.post('/new', JSON.stringify($scope.new), postConfig).then(function(){
            console.log('success');
        });
    };
});

function initNewPost(){
    return {text: '', title: ''};
}