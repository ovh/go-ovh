package ovh

import (
	"net/http"
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func TestErrorString(t *testing.T) {
	err := &APIError{
		Code:    http.StatusBadRequest,
		Message: "Bad request",
	}
	td.CmpString(t, err, `HTTP Error 400: "Bad request"`)

	err = &APIError{
		Code:    http.StatusConflict,
		Message: `the cart id "foobar" already exists`,
		Class:   "CartAlreadyExists",
		QueryID: "EU.ext-99.foobar",
	}
	td.CmpString(t, err, `HTTP Error 409: CartAlreadyExists: "the cart id \"foobar\" already exists" (X-OVH-Query-Id: EU.ext-99.foobar)`)
}
