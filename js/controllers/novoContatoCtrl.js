angular.module("listaTelefonica").controller("novoContatoCtrl", function ($scope, serialGenerator, $location, operadorasAPI, flashService) {
	
	$scope.adicionarContato = function (contato) {
		contato.serial = serialGenerator.generate();
		contatosAPI.saveContato(contato).success(function (data) {
			delete $scope.contato;
			$scope.contatoForm.$setPristine();
			$location.path("/contatos");
		});
	};
	$scope.contatoRedirect = function () {
		$location.path("/contatos");
	};

	var buscarOperadoras = function (){
		operadorasAPI.getOperadoras()
		.then(function(response){
			$scope.operadoras = response.data;
			flashService.success('Operadoras carregas');
		})
		.catch(function(response){			
			flashService.error('Erro: ' + response.statusText);
		})
	}

	buscarOperadoras();
});