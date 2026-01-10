package router

import "net/http"

type Router struct {
	handler http.Handler
}

func New() *Router {
	return &Router{
		handler: http.NewServeMux(),
	}
}

func (r *Router) Handle(pattern string, hander http.Handler) {
	r.handler.(*http.ServeMux).Handle(pattern, hander)
}

func (r *Router) Use(mv func(http.Handler) http.Handler) {
	r.handler = mv(r.handler)
}

func (r *Router) Handler() http.Handler {
	return r.handler
}
