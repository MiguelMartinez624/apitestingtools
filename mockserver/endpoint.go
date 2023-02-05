package mockserver

type Endpoint[T interface{}] struct {
	path    string
	method  string
	handler T
}

func NewEndpoint[T interface{}](method, path string, handler T) Endpoint[T] {
	return Endpoint[T]{path: path, method: method, handler: handler}
}
