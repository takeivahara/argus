package route

import (
	"net/http"
	"github.com/gorilla/mux"
	"argus/controller"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
		Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"HelloName",
		"GET",
		"/hello/{name}",
		controller.HelloName,
	},

	Route{
		"Operadoras",
		"GET",
		"/operadoras",
		controller.Operadoras,
	},

	Route{
		"Users",
		"GET",
		"/users",
		controller.Users,
	},

	Route{
		"Users",
		"OPTIONS",
		"/users",
		controller.Users,
	},

	Route{
		"Contato",
		"GET",
		"/contato",
		controller.Contato,
	},

	Route{
		"BuscarTodasPessoas",
		"GET",
		"/buscarTodasPessoas",
		controller.BuscarTodasPessoa,
	},

	Route{
		"Concorrencia",
		"GET",
		"/concorrencia",
		controller.Concorrencia,
	},
}

