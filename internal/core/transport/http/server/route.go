package core_http_server

import (
	"net/http"
)

type Route struct {
	Method 	string
	Path   	string // request's url 
	Handler http.HandlerFunc
}

func NewRoute(method, path string, handler http.HandlerFunc) Route {
	return Route{
		Method:  method,
		Path: 	 path,
		Handler: handler,
	}
}