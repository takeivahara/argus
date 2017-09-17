angular.module("listaTelefonica").controller("loginCtrl", function ($scope, $location, authenticationService, flashService) {   

    $scope.login = function (user) {    
        authenticationService.Login(user.username, user.password, function (response) {
            if (response.success) {
                authenticationService.SetCredentials(user.username, user.password);
                $location.path('/contatos');
            } else {
                flashService.error(response.message);
            }
        });
    };

    (function initController() {        
        // reset login status
        authenticationService.ClearCredentials();
    })();
});