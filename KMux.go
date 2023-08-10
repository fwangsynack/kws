package kws

import "net/http"

type KMux struct {
	middlewares []func(handler http.Handler) http.Handler

	handler *http.ServeMux
}

func NewKMux() *KMux {
	mx := &KMux{
		handler: http.NewServeMux(),
	}
	return mx
}

func (mx *KMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mx.handler.ServeHTTP(w, r)
}

func (mx *KMux) Use(middlewares ...func(http.Handler) http.Handler) {
	if mx.handler == nil {
		panic("KMux handler is nil")
	}
	mx.middlewares = append(mx.middlewares, middlewares...)
}

func (mx *KMux) wrap(handler http.Handler) http.Handler {
	middlewares := mx.middlewares
	if len(middlewares) == 0 {
		return handler
	}

	h := middlewares[len(middlewares)-1](handler)
	for i := len(middlewares) - 2; i >= 0; i-- {
		h = middlewares[i](h)
	}

	return h
}

func (mx *KMux) Handle(pattern string, handler http.Handler) {
	h := mx.wrap(handler)
	mx.handler.Handle(pattern, h)
}

func (mx *KMux) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	mx.Handle(pattern, http.HandlerFunc(handler))
}
