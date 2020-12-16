package routes

import (
	"net/http"

	"restful/response"
	"github.com/gorilla/mux"
)

type Route struct {
	Method     string
	Path    string
	Handler    http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

var routes []Route

func init() {
	//apiMethodbind("GET", "/api/root", response.FindAllApi, nil)
	apiMethodbind("GET", "/api/movies/id/{id}", response.FindMovieById, nil)
	apiMethodbind("GET", "/api/movies/name/{name}", response.FindMovieByName, nil)
	apiMethodbind("GET", "/api/movies", response.FindAllMovies, nil)
	apiMethodbind("POST", "/api/movies", response.CreateMovie, nil)
	apiMethodbind("PUT", "/api/movies", response.UpdateMovie, nil)
	apiMethodbind("DELETE", "/api/movies/{id}", response.DeleteMovieById, nil)
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	for _, route := range routes {
		r.Methods(route.Method).Path(route.Path).Handler(route.Handler)
		if route.Middleware != nil {
			r.Use(route.Middleware)
		}
	}
	return r
}

func apiMethodbind(method, path string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, Route{method, path, handler, middleware})
}
