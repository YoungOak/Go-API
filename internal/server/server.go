package server

import (
	"net/http"
)

type Router interface {
	AddHandler(route string, handler http.HandlerFunc)
	Serve() error
}

type router struct {
	address string
	router  *http.ServeMux
}

func NewRouter(listenAddress string) Router {
	return &router{
		listenAddress,
		http.NewServeMux(),
	}
}

func (r *router) AddHandler(route string, handler http.HandlerFunc) {
	r.router.HandleFunc(route, handler)
}

func (r *router) Serve() error {

	server := http.Server{
		Addr:    r.address,
		Handler: r.router,
	}

	return server.ListenAndServe()
}
