package core

import (
	"net/http"
	"regexp"
	"strings"
)

type Middleware func(http.ResponseWriter, *http.Request, func(http.ResponseWriter, *http.Request))

type Route struct {
	Method     string
	Pattern    string
	Handler    func(http.ResponseWriter, *http.Request, map[string]string)
	Name       string
	Middleware []Middleware
}

type Router struct {
	routes           []Route
	basePath         string
	groups           []string
	globalMiddleware []Middleware
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Use(middleware Middleware) {
	r.globalMiddleware = append(r.globalMiddleware, middleware)
}

func (r *Router) AddRoute(method, pattern string, handler func(http.ResponseWriter, *http.Request, map[string]string), name string, middlewares ...Middleware) {
	fullPattern := r.basePath + pattern
	fullPattern = strings.ReplaceAll(fullPattern, "{", "(?P<")
	fullPattern = strings.ReplaceAll(fullPattern, "}", ">[^/]+)")
	r.routes = append(r.routes, Route{
		Method:     method,
		Pattern:    "^" + fullPattern + "$",
		Handler:    handler,
		Name:       name,
		Middleware: middlewares,
	})
}

func (r *Router) applyMiddlewareChain(middlewares []Middleware, handler http.HandlerFunc) http.HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		middleware := middlewares[i]
		next := handler
		handler = func(w http.ResponseWriter, req *http.Request) {
			middleware(w, req, next)
		}
	}
	return handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if req.Method == route.Method {
			regex := regexp.MustCompile(route.Pattern)
			matches := regex.FindStringSubmatch(req.URL.Path)

			if matches != nil {
				// Extract named parameters
				params := map[string]string{}
				for i, name := range regex.SubexpNames() {
					if i > 0 && name != "" {
						params[name] = matches[i]
					}
				}

				// Apply middleware and handler
				handler := func(w http.ResponseWriter, r *http.Request) {
					route.Handler(w, r, params)
				}

				middlewareChain := append(r.globalMiddleware, route.Middleware...)
				r.applyMiddlewareChain(middlewareChain, handler)(w, req)
				return
			}
		}
	}
	http.NotFound(w, req)
}
