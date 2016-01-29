/**
 * Created by serhii on 1/29/16.
 */

app.controller('OneCtr', function ($scope, $http) {

    $scope.inEditor = '';

    $scope.postNew = function () {

        $scope.knot.raw = editor.value();
        var knot = angular.copy($scope.knot);
        console.log(knot);

        if (knot.id) {
            $http.put('/knots/' + knot.id, JSON.stringify(knot), postConfig).then(function () {
                $scope.router('one', knot.id);
            });
        } else {
            $http.post('/knots', JSON.stringify(knot), postConfig).then(function () {
                $('#editor').modal('hide');
                $scope.router('main');
            });
        }

    };

    $scope.remove = function (type, knot) {
        if (confirm("Are you sure want to delete?")) {
            $http.delete('/knots/' + knot.id, {}).then(function () {
                $scope.router('main');
            });
        }
    };

    $scope.edit = function (type, knot) {
        $scope.editorAtion = "update";
        $scope.knot = knot;
        if ($scope.inEditor != knot.id){
            editor.value(knot.raw);
            $scope.inEditor = knot.id;
        }
    };


});

