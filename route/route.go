package route

import (
	controller "github.com/wei840222/todo-api/controller"

	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Method     string
	Pattern    string
	Handler    http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

var routes []Route

func init() {
	register("GET", "/todo", controller.AllTodo, nil)
	register("GET", "/todo/{id}", controller.FindTodo, nil)
	register("POST", "/todo", controller.CreateTodo, nil)
	register("PUT", "/todo", controller.UpdateTodo, nil)
	register("DELETE", "/todo/{id}", controller.DeleteTodo, nil)
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	for _, route := range routes {
		r := router.Methods(route.Method).
			Path(route.Pattern)
		if route.Middleware != nil {
			r.Handler(route.Middleware(route.Handler))
		} else {
			r.Handler(route.Handler)
		}
	}
	return router
}

func register(method string, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, Route{method, pattern, handler, middleware})
}
