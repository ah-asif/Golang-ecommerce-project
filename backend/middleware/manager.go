package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}

}

func (mngr *Manager) Use(middleware ...Middleware) *Manager {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middleware...)
	return mngr
}

func (mngr *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {
	h := handler

	// [Logger, LoggerHudai, CorsWithPrefligt]
	// h = CorsWithPrefligt(LoggerHudai(Logger(mux)))
	for _, middleware := range middlewares {
		h = middleware(h)
	}

	return h
}

func (mngr *Manager) WrapMux(handler http.Handler) http.Handler {
	h := handler

	// [Cors, Preflight,Logger]
	// h = Logger(Preflight(Cors(mux)))
	for _, middleware := range mngr.globalMiddlewares {
		h = middleware(h)
	}

	return h
}
