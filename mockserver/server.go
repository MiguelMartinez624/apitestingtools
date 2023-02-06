package mockserver

import (
	"net/http"
	"net/http/httptest"
)

type HTTPServer[T interface{}] interface {
	SetEndpoint(path, method string, handlers []T)
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

type Server[T interface{}] struct {
	routes map[string]Endpoint[T]
	runner HTTPServer[T]
}

func (s *Server[T]) UseServer(runner HTTPServer[T]) *Server[T] {
	s.runner = runner
	return s
}

func (s *Server[T]) Endpoint(method, path string, handler ...T) *Server[T] {
	if _, ok := s.routes[path]; ok {
		panic("route already exist")
	}

	endpoint := NewEndpoint[T](path, method, handler)
	s.routes[path] = endpoint

	s.runner.SetEndpoint(path, method, handler)

	return s
}

func (s *Server[T]) SendRequest(request *http.Request) *RequestResult {
	w := httptest.NewRecorder()
	s.runner.ServeHTTP(w, request)

	return &RequestResult{
		output: &RequestResponse{
			Code:    w.Code,
			Date:    nil,
			BodyStr: w.Body.String(),
			Body:    w.Body,
		},
	}
}

func NewTestServer[T interface{}]() *Server[T] {
	return &Server[T]{
		routes: make(map[string]Endpoint[T], 0),
		runner: nil,
	}
}
