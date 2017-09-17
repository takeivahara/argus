angular.module("listaTelefonica").factory("contatosAPI", function ($http, config) {
	
	delete $http.defaults.headers.common['X-Requested-With'];

	var _getContatos = function () {
		return $http.get(config.baseUrl + "/contatos")


				// $http.get(baseUrl).success(function(response) {
				//   list = response.data;
				// }).error(function(err) {
				//   console.log(err);
				// });
  		
	};

	var _getContato = function (id) {
		return $http.get(config.baseUrl + "/contatos/" + id);
	};

	var _saveContato = function (contato) {
		return $http.post(config.baseUrl + "/contatos", contato);
	};

	return {
		getContatos: _getContatos,
		getContato: _getContato,
		saveContato: _saveContato
	};
});