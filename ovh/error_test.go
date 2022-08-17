package ovh

import (
	"fmt"
	"net/http"
	"testing"
)

func TestErrorString(t *testing.T) {
	err := &APIError{
		Code:    http.StatusBadRequest,
		Message: "Bad request",
	}

	expected := `HTTP Error 400: "Bad request"`
	got := fmt.Sprintf("%s", err)

	if got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}

	err.Class = "CartAlreadyExists"
	err.Code = http.StatusConflict
	err.Message = `the cart id "foobar" already exists`
	err.QueryID = "EU.ext-99.foobar"

	expected = `HTTP Error 409: CartAlreadyExists: "the cart id \"foobar\" already exists" (X-OVH-Query-Id: EU.ext-99.foobar)`
	got = fmt.Sprintf("%s", err)

	if got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}
