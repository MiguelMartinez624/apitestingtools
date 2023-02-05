package mockserver

import (
	"bytes"
	"testing"
)

type RequestResponse struct {
	Code    int
	Date    interface{}
	BodyStr string
	Body    *bytes.Buffer
}

type RequestResult struct {
	output *RequestResponse
}

// here we do the assertions.

func (receiver RequestResult) Asserts(t *testing.T, asserts ...AssertFunc) {
	for _, assert := range asserts {
		assert(t, receiver.output)
	}
}
