package router

import (
	"github.com/gorilla/mux"
	"github.com/yangyouwei/golangapi/handler"
	"github.com/yangyouwei/golangapi/log"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{Name:"Index",     Method:"GET",   Pattern:"/",                HandlerFunc:handler.Index},
	Route{Name:"TodoIndex", Method:"GET",   Pattern:"/todos",           HandlerFunc:handler.TodoIndex},
	Route{Name:"TodoShow",  Method:"GET",   Pattern:"/todos/{todoId}",  HandlerFunc:handler.TodoShow},
}



func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = log.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}