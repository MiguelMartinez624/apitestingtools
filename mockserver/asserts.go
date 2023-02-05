package mockserver

import (
	"encoding/json"
	"testing"
)

type AssertFunc func(t *testing.T, response *RequestResponse)

func IsStatus(status int) AssertFunc {
	return func(t *testing.T, response *RequestResponse) {
		if response.Code != status {
			t.Errorf("expected [%v] got [%v]", status, response.Code)
		}
	}
}

func StringBodyEquals(expectedStr string) AssertFunc {
	return func(t *testing.T, response *RequestResponse) {
		if response.BodyStr != expectedStr {
			t.Errorf("expected: \n %v \n got \n %v", expectedStr, response.BodyStr)
		}
	}
}

func Equals(ev func(responseBody string) bool, errorMessage string) AssertFunc {

	return func(t *testing.T, response *RequestResponse) {
		pass := ev(response.BodyStr)
		if !pass {
			t.Error(errorMessage)
		}
	}
}

func AsStruct[T interface{}](ev func(t *testing.T, responseBody T)) AssertFunc {

	return func(t *testing.T, response *RequestResponse) {
		var result T
		if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
			panic(err)
		}
		ev(t, result)

	}
}
