angular.module("listaTelefonica").config(function ($routeProvider, $locationProvider) {	

	//$locationProvider.html5Mode(true);

	$routeProvider.when("/", {
		templateUrl: "view/login.html",
		controller: "loginCtrl"
	});
	$routeProvider.when("/login", {
		templateUrl: "view/login.html",
		controller: "loginCtrl"
	});
	$routeProvider.when("/contatos", {
		templateUrl: "view/contatos.html",
		controller: "contatosCtrl",
		resolve: {
			contatos: function (contatosAPI) {				
				return contatosAPI.getContatos();
			}
		}		
	});
	$routeProvider.when("/testeControles", {
		templateUrl: "view/testeControles.html",
		controller: "testeControlesCtrl"
	});
	$routeProvider.when("/novoContato", {
		templateUrl: "view/novoContato.html",
		controller: "novoContatoCtrl"
	});
	$routeProvider.when("/detalhesContato/:id", {
		templateUrl: "view/detalhesContato.html",
		controller: "detalhesContatoCtrl",
		resolve: {
			contato: function (contatosAPI, $route) {
				return contatosAPI.getContato($route.current.params.id);
			}
		}
	});
	$routeProvider.when("/error", {
		templateUrl: "view/error.html"
	});
	$routeProvider.otherwise({redirectTo: "/"});
});