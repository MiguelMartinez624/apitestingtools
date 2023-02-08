package mockserver

import (
	"bytes"
	"encoding/json"
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

func (receiver RequestResult) Result(output interface{}) {
	if err := json.NewDecoder(receiver.output.Body).Decode(output); err != nil {
		panic(err)
	}
}

func (receiver RequestResult) Code() int {
	return receiver.output.Code
}

func (receiver RequestResult) String() string {
	return receiver.output.BodyStr
}
