(function () {
    'use strict';

    angular
        .module('listaTelefonica')
        .factory('userAPIService', userAPIService);

    userAPIService.$inject = ['$http', 'config'];
    function userAPIService($http, config) {
        var service = {};

        service.GetAll = GetAll;
        service.GetById = GetById;
        service.GetByUsername = GetByUsername;
        service.Create = Create;
        service.Update = Update;
        service.Delete = Delete;

        return service;

        function GetAll() {
            console.log(config.baseUrl + "/users");
            return $http.get(config.baseUrl + "/users").then(handleSuccess, handleError('Error getting all users'));            
        }

        function GetById(id) {
            return $http.get(config.baseUrl + '/users/' + id).then(handleSuccess, handleError('Error getting user by id'));
        }

        function GetByUsername(username) {
            return $http.get(config.baseUrl + '/users/' + username).then(handleSuccess, handleError('Error getting user by username'));
        }

        function Create(user) {
            return $http.post(config.baseUrl + '/users', user).then(handleSuccess, handleError('Error creating user'));
        }

        function Update(user) {
            return $http.put(config.baseUrl + '/users/' + user.id, user).then(handleSuccess, handleError('Error updating user'));
        }

        function Delete(id) {
            return $http.delete(config.baseUrl + '/users/' + id).then(handleSuccess, handleError('Error deleting user'));
        }

        // private functions

        function handleSuccess(res) {
            return res.data;
        }

        function handleError(error) {
            return function () {
                return { success: false, message: error };
            };
        }
    }

})();
