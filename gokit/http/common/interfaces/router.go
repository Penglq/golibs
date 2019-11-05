package interfaces

import "github.com/gorilla/mux"

type Routerfunc func(*mux.Router)

func RouterOptions(r *mux.Router, options ...Routerfunc) {
	for _, router := range options {
		router(r)
	}
}

type SubRouter interface {
	New()
}

type SubRouterfunc func(*mux.Router)

func SubRouterOptions(r *mux.Router, options ...SubRouterfunc) {
	for _, subRouter := range options {
		subRouter(r)
	}
}
