/**
 * Created by serhii on 1/29/16.
 */

app.controller('RouterCtr', function ($scope, $http, $location, Tools) {
    $scope.Tools = Tools;

    $scope.router = function(action, id){
        $scope.routeAction = action;
        console.log([action, id]);

        if (action == '404'){
            $location.path('404');
        }
        if(action == 'one'){
            $location.path('one/'+id);

            if(id == 'new'){
                initEmptyKnot($scope);

            } else {
                $http.get("/knots/" + id).then(function (response) {
                    $scope.knot = response.data.knot;
                }, function() {
                    $scope.router('404');
                });
            }

        }

        if(action == 'main'){
            $location.path("");
            $scope.knots = [];
            getKnots($http, $scope);
        }
    };

    $scope.isRoute = function (route) {
        return route == $scope.routeAction;
    };

    initPage($scope, $location);

    $scope.empty = function () {
        $scope.editorAtion = "new";
        initEmptyKnot($scope);
        editor.value("");

    };
});

function initPage($scope, $location){
    var path = $location.path();

    if(path == "" || path == '/') {
        $scope.router('main');
    } else {
        if(path.split('/').length == 1){
            $scope.router(path.split('/')[1]);
        }else{
            $scope.router(path.split('/')[1], path.split('/')[2]);
        }
    }
}

function getKnots($http, $scope) {
    $http.get("/knots").then(function (response) {
        $scope.knots = response.data.knots;
    });
}

function initEmptyKnot($scope) {
    $scope.knot = {raw: '', title: ''};
}