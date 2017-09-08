angular.module("listaTelefonica", ["ngMessages", "serialGenerator", "ui", "ngRoute", "ngCookies", "ui.directives"]);

angular.module("listaTelefonica").run(function ($rootScope, $location, $cookies, $http) {

	// keep user logged in after page refresh
    $rootScope.globals = $cookies.getObject('globals') || {};
    
    if ($rootScope.globals.currentUser) {
        $http.defaults.headers.common['Authorization'] = 'Basic ' + $rootScope.globals.currentUser.authdata;
    }

    $rootScope.$on('$locationChangeStart', function (event, next, current) {        

        // redirect to login page if not logged in and trying to access a restricted page
        var restrictedPage = $.inArray($location.path(), ['/login']) === -1;
        var loggedIn = $rootScope.globals.currentUser;

        if (restrictedPage && !loggedIn) {            
            $location.path('/login');
        }
    });
});