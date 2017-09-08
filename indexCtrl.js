angular.module("listaTelefonica").controller("indexCtrl", function ($scope, $location, authenticationService) {
	
	$scope.menuClick = function (destino) {
		$location.path("/" + destino);
	};

	$scope.logout = function () {		
		authenticationService.ClearCredentials();
		$location.path("/login");
	};
});